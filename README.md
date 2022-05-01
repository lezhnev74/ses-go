# PoC: Event Set Pattern Matching (in Go)

Regular expressions are a notation for describing sets of character strings, this package offers expressions for sets of events.
**The idea is that an analyst can express complex behaviour in terms of facts (events) and find matching actors by scanning a stream of events**.

Note: this is a proof of concept, written in Go.

## Concept Background

I have always been interested in regular expressions and in general pattern matching. At work, I program systems that emit a lot of data, including a vast amount of events.
At some point I came across an interesting paper called [Sequenced event set pattern matching](https://dl.acm.org/doi/10.1145/1951365.1951372)
and it's later revision [Efficient event pattern matching with match windows](https://dl.acm.org/doi/10.1145/2339530.2339607).
Ideas from the papers sparked my imagination. It was tempting to make a language for recognizing events patterns in a stream. So I did.

Similar ideas are used in marketing, see [Segment Personas](https://segment.com/product/personas/), as well as in databases
see [Vertica Match](https://www.vertica.com/docs/9.2.x/HTML/Content/Authoring/SQLReferenceManual/Statements/SELECT/MATCHClause.htm)
and [Snowflake Match Recognize](https://docs.snowflake.com/en/sql-reference/constructs/match_recognize.html).

## Query Language

```
// Sample Query: find all web sessions where users read blog before purchasing

within 3 days                                    // set a window for a match
     event website_visit+ where page~="blog/.*"  // find at least one blog visit
then event website_visit+ where page="cart.html" // then find at least one cart visit
then within 1 hour event purchase                // and then there must be a purchase within 1 hour after the cart visit
group by session_id                              // so all events must be within one session
```

It is an SQL-like declarative language that should be familiar to anyone who works with data. Combine that with regular expression syntax
and here you go. It is quite simple to express any complex sequential facts.

## Technology Details

This program can work with any data source. It needs an adapter that converts data coming from a given data source (apache kafka, monetdb, etc)
to the format that this program understands. See `automaton.Event` interface.

Since the data source is abstracted, this system does not know the schema of data, so it can't apply static typing. Instead,
it follows a more flexible way of dynamic typing and checks types at run-time.

For language parsing I used [Antlr](https://www.antlr.org/) which is just awesome :)

## Usage Example

Run `TestPublicExample()` in `automaton_test.go`.

```go
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
    runner := MakeRunner(ses)
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
```