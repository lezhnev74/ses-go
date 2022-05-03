package automaton

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"

	"ses_pm_antlr/automaton/state"
	"ses_pm_antlr/parser"
)

func TestFailedAutomatons(t *testing.T) {
	tests := []string{
		// qty underflow
		`event visit{4,} where section="content" group by session`, // there is 3 "content"
		// qty overflow
		`event visit where section="content"
		 event visit{0} where section="shop" // there 1+ "shop"
		 group by session`,
		// no match
		`event unknown group by id`, // there is 3 "content"
	}

	// parse JSON events
	var events []AnyData
	jsonEvents, err := os.ReadFile("_testdata/visits.json")
	if err != nil {
		t.Fatalf("unable to open file: %s", err)
	}
	err = json.Unmarshal(jsonEvents, &events)
	if err != nil {
		t.Fatalf("invalid json text: %s", err)
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			db := state.MakeBadgerDb(testName, fmt.Sprintf("%s/%s", t.TempDir(), testName))
			ses := parser.ParseSESQuery(tt)
			runner := MakeRunner(ses, db)

			// feed to the automaton
			for _, event := range events {
				e := &SimpleEvent{event}
				runner.Accept(e)
			}

			// assert the state of the automaton at the end of events
			if len(runner.GetAcceptingAutomatons()) != 0 {
				t.Errorf("No accepting automatons should be found")
			}
		})
	}
}

func TestDeterministicAutomaton(t *testing.T) {
	type testData struct {
		query            string
		expectedEventIds []map[string][]any // result of matching
		expectedGroupBy  any
	}

	tests := []testData{
		// find a session of successful purchase
		{
			`event visit{2,} where section="content"
	    then event visit where section="shop" and page~="products.*"
		then event visit where section="shop" and page="checkout.html"
		then event visit where section="shop" and page="thank-you.html" // reached this page
		     group by session
            `,
			[]map[string][]any{
				0: {"visit": {"wd1", "h4w"}},
				1: {"visit": {"eh3"}},
				2: {"visit": {"43g"}},
				3: {"visit": {"3h2"}},
			},
			"s1",
		},
		// find a session of interrupted purchase
		{
			`event visit{3,} where section="content"
		then event visit  where section="shop" and page~="products.*"
		then event visit  where section="shop" and page="checkout.html"
		then event visit{0}  where section="shop" and page="thank-you.html" // never reached this page
		    group by session
		   `,
			[]map[string][]any{
				0: {"visit": {"s4g", "b23", "na9"}},
				1: {"visit": {"0gk"}},
				2: {"visit": {"g22"}},
			},
			"s2",
		},
	}

	// parse JSON events
	var events []AnyData
	jsonEvents, err := os.ReadFile("_testdata/visits.json")
	if err != nil {
		t.Fatalf("unable to open file: %s", err)
	}
	err = json.Unmarshal(jsonEvents, &events)
	if err != nil {
		t.Fatalf("invalid json text: %s", err)
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			db := state.MakeBadgerDb(testName, fmt.Sprintf("%s/%s", t.TempDir(), testName))
			ses := parser.ParseSESQuery(tt.query)
			runner := MakeRunner(ses, db)

			// feed to the automaton
			for _, event := range events {
				e := &SimpleEvent{event}
				runner.Accept(e)
			}

			// assert the state of the automaton at the end of events
			acceptingAutomatons := runner.GetAcceptingAutomatons()
			if l := len(acceptingAutomatons); l != 1 {
				t.Fatalf("one accepting automaton expected, found %d", l)
			}
			automaton := acceptingAutomatons[0] // pick the only one
			actualEventIds := automaton.GetAcceptedEventIdsAsSlice()

			if !reflect.DeepEqual(actualEventIds, tt.expectedEventIds) {
				t.Errorf("automaton accepted events\n%v(%[1]T)\ndo not match expected\n%v(%[2]T)\n", actualEventIds, tt.expectedEventIds)
			}
			if !reflect.DeepEqual(automaton.groupBy, tt.expectedGroupBy) {
				t.Errorf("automaton groupby %v(%[1]T) does not match expected %v(%[2]T)\n", automaton.groupBy, tt.expectedGroupBy)
			}
		})
	}
}

func TestPublicExample(t *testing.T) {
	jsonData := `
[
  {"id": "1", "name": "website_visit", "page": "home.html", "session_id": "s1"},
  {"id": "2", "name": "website_visit", "page": "blog/post1.html", "session_id": "s1"},
  {"id": "3", "name": "website_visit", "page": "blog/post2.html", "session_id": "s1"},
  {"id": "4", "name": "website_visit", "page": "cart.html", "session_id": "s1"},
  {"id": "5", "name": "purchase", "session_id": "s1"}
]
`

	query := `
// Sample Query: find all web sessions where users read blog before purchasing

     event website_visit+ where page~="blog/.*"  
then event website_visit+ where page="cart.html" 
then event purchase                
group by session_id                	
`

	// Read events from JSON
	var events []AnyData
	err := json.Unmarshal([]byte(jsonData), &events)
	if err != nil {
		log.Fatalf("invalid json text: %s", err)
	}

	// Parse Query
	ses := parser.ParseSESQuery(query)

	// Run Matching
	db := state.MakeBadgerDb("scope1", fmt.Sprintf("%s/%s", t.TempDir(), "scope1"))
	runner := MakeRunner(ses, db)
	for _, event := range events {
		e := &SimpleEvent{event} // wrap your data to something that implement automaton.Event interface
		runner.Accept(e)
	}

	// Find matching events
	// Automatons that are in accepting state contain matching events
	if len(runner.GetAcceptingAutomatons()) == 0 {
		log.Fatalf("No accepting automatons found")
	}

	for _, a := range runner.GetAcceptingAutomatons() {
		fmt.Printf("Matching session: %s\n", a.groupBy) // duplicates are possible
	}
}
