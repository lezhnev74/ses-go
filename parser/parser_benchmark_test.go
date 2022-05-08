package parser

import (
	"testing"
	"time"
)

func BenchmarkQueryParse(b *testing.B) {
	q := `
	within 3 days                                // set a window for a match
     event website_visit+ where page~="blog/.*"  // find at least one blog visit
then event website_visit+ where page="cart.html" // then find at least one cart visit
then within 1 hour event purchase                // and then there must be a purchase within 1 hour after the cart visit
group by session_id                              // so all events must be within one session
`

	for i := 0; i < b.N; i++ {
		ParseSESQuery(q, time.Now())
	}
}
