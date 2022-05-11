package automaton

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"ses_pm_antlr/automaton/state"
	"ses_pm_antlr/parser"
	"ses_pm_antlr/vector"
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
			ses := parser.ParseSESQuery(tt, time.Now())
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

func TestSetWindow(t *testing.T) {
	now := time.Now()

	type testInput struct {
		query          string
		events         []AnyData
		matched_groups [][]any
	}

	tests := []testInput{
		{
			`event a then within 1 minute event b 
			 group by session`,
			[]AnyData{
				// b fits within the window
				{"id": "1", "name": "a", "time": now.UnixNano(), "session": "s1"},
				{"id": "2", "name": "b", "time": now.Add(time.Second).UnixNano(), "session": "s1"},
				// b is outside the window
				{"id": "3", "name": "a", "time": now.UnixNano(), "session": "s2"},
				{"id": "4", "name": "b", "time": now.Add(2 * time.Minute).UnixNano(), "session": "s2"},
			},
			[][]any{{"s1"}},
		},
		{
			`event a then skip 2 seconds and within 2 seconds event b 
			 group by session`,
			[]AnyData{
				// b fits within the window
				{"id": "1", "name": "a", "time": now.UnixNano(), "session": "s1"},
				{"id": "2", "name": "b", "time": now.Add(3 * time.Second).UnixNano(), "session": "s1"},
				// b is outside the window
				{"id": "3", "name": "a", "time": now.UnixNano(), "session": "s2"},
				{"id": "4", "name": "b", "time": now.Add(1 * time.Second).UnixNano(), "session": "s2"},
			},
			[][]any{{"s1"}},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("tt %d", i), func(t *testing.T) {
			db := state.MakeBadgerDb("scope1", "")
			ses := parser.ParseSESQuery(tt.query, time.Now())
			runner := MakeRunner(ses, db)

			for _, e := range tt.events {
				runner.Accept(&SimpleEvent{e})
			}

			sessions := [][]any{}
			for _, a := range runner.GetAcceptingAutomatons() {
				sessions = append(sessions, a.groupBy)
			}

			// remove duplicates
			for i := 0; i < len(sessions); i++ {
				for j := i + 1; j < len(sessions); j++ {
					if vector.CmpAnySlices(sessions[i], sessions[j]) {
						sessions = append(sessions[:j], sessions[j+1:]...)
					}
				}
			}

			if !reflect.DeepEqual(sessions, tt.matched_groups) {
				t.Errorf("Unexpected sessions found %v, expected %v", sessions, tt.matched_groups)
			}
		})
	}
}

func TestSesWindow(t *testing.T) {
	now := time.Now()
	tests := []struct {
		query          string
		events         []AnyData
		matched_groups [][]any
	}{
		{
			"within 1 second event a and then event a group by session",
			[]AnyData{
				{"id": "1", "name": "a", "time": now.UnixNano(), "session": "s1"},
				{"id": "2", "name": "a", "time": now.Add(2 * time.Second).UnixNano(), "session": "s1"},
				{"id": "3", "name": "a", "time": now.Add(3 * time.Second).UnixNano(), "session": "s1"},
			},
			[][]any{{"s1"}},
		},
		{
			"window within 5 seconds; event a+ and then event a{2,} group by session",
			[]AnyData{
				{"id": "1", "name": "a", "time": now.UnixNano(), "session": "s1"},
				{"id": "2", "name": "a", "time": now.Add(3 * time.Second).UnixNano(), "session": "s1"},
				{"id": "3", "name": "a", "time": now.Add(6 * time.Second).UnixNano(), "session": "s1"},
				{"id": "4", "name": "a", "time": now.Add(7 * time.Second).UnixNano(), "session": "s1"},
			},
			[][]any{{"s1"}},
		},
		{
			"window within 1 second; event a and then event b group by session",
			[]AnyData{
				// two events match the window
				{"id": "1", "name": "a", "time": now.UnixNano(), "session": "s1"},
				{"id": "2", "name": "b", "time": now.Add(time.Second).UnixNano(), "session": "s1"},
				// events are too far, outside the window
				{"id": "3", "name": "a", "time": now.UnixNano(), "session": "s2"},
				{"id": "4", "name": "b", "time": now.Add(2 * time.Second).UnixNano(), "session": "s2"},
			},
			[][]any{{"s1"}},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("tt %d", i), func(t *testing.T) {
			db := state.MakeBadgerDb(fmt.Sprintf("tt %d", i), "")
			ses := parser.ParseSESQuery(tt.query, time.Now())
			runner := MakeRunner(ses, db)

			for _, e := range tt.events {
				runner.Accept(&SimpleEvent{e})
			}

			sessions := [][]any{}
			for _, a := range runner.GetAcceptingAutomatons() {
				sessions = append(sessions, a.groupBy)
			}

			// remove duplicates
			for i := 0; i < len(sessions); i++ {
				for j := i + 1; j < len(sessions); j++ {
					if vector.CmpAnySlices(sessions[i], sessions[j]) {
						sessions = append(sessions[:j], sessions[j+1:]...)
					}
				}
			}

			if !reflect.DeepEqual(sessions, tt.matched_groups) {
				t.Errorf("Unexpected sessions found %v, expected %v", sessions, tt.matched_groups)
			}
		})
	}
}

func TestVariousAspects(t *testing.T) {
	now := time.Now()

	type testInput struct {
		query          string
		events         []AnyData
		matched_groups [][]any
	}

	tests := []testInput{
		{
			`event a+ then event b where id=any(a.id) 
			 group by session`,
			[]AnyData{
				{"id": "1", "name": "a", "time": now.UnixNano(), "session": "s1"},
				{"id": "2", "name": "a", "time": now.Add(1 * time.Nanosecond).UnixNano(), "session": "s1"},
				{"id": "3", "name": "a", "time": now.Add(2 * time.Nanosecond).UnixNano(), "session": "s1"},
				{"id": "2", "name": "b", "time": now.Add(3 * time.Nanosecond).UnixNano(), "session": "s1"},
			},
			[][]any{{"s1"}},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("tt %d", i), func(t *testing.T) {
			db := state.MakeBadgerDb("scope1", "")
			ses := parser.ParseSESQuery(tt.query, time.Now())
			runner := MakeRunner(ses, db)

			for _, e := range tt.events {
				runner.Accept(&SimpleEvent{e})
			}

			sessions := [][]any{}
			for _, a := range runner.GetAcceptingAutomatons() {
				sessions = append(sessions, a.groupBy)
			}

			// remove duplicates
			for i := 0; i < len(sessions); i++ {
				for j := i + 1; j < len(sessions); j++ {
					if vector.CmpAnySlices(sessions[i], sessions[j]) {
						sessions = append(sessions[:j], sessions[j+1:]...)
					}
				}
			}

			if !reflect.DeepEqual(sessions, tt.matched_groups) {
				t.Errorf("Unexpected sessions found %v, expected %v", sessions, tt.matched_groups)
			}
		})
	}
}

func TestAutomatonStaticDataSet(t *testing.T) {
	type testData struct {
		query            string
		expectedEventIds []map[string][]any // result of matching
		expectedGroupBy  []any
	}

	tests := []testData{
		// complex where-expr
		{
			`event visit{5,} where (section="content" or section="shop") and page!="careers.html"
		     group by session
            `,
			[]map[string][]any{
				0: {"visit": {"wd1", "h4w", "eh3", "43g", "3h2"}},
			},
			[]any{"s1"},
		},
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
			[]any{"s1"},
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
			[]any{"s2"},
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
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			testName := t.Name()
			db := state.MakeBadgerDb(testName, fmt.Sprintf("%s/%s", t.TempDir(), testName))
			ses := parser.ParseSESQuery(tt.query, time.Now())
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
  {"id": "1", "name": "website_visit", "page": "home.html", "session_id": "s1", "time":1},
  {"id": "2", "name": "website_visit", "page": "blog/post1.html", "session_id": "s1", "time":2},
  {"id": "3", "name": "website_visit", "page": "blog/post2.html", "session_id": "s1", "time":3},
  {"id": "4", "name": "website_visit", "page": "cart.html", "session_id": "s1", "time":4},
  {"id": "5", "name": "purchase", "session_id": "s1", "time":5}
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
	ses := parser.ParseSESQuery(query, time.Now())

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
