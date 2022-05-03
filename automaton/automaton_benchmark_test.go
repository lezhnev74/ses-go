package automaton

import (
	"fmt"
	"testing"
	"time"

	"ses_pm_antlr/automaton/state"
	"ses_pm_antlr/parser"
)

func BenchmarkRunnerMemoryMode(b *testing.B) {
	db := state.MakeBadgerDb("scope1", "")
	benchmarkRunner(b, db)
}

func BenchmarkRunnerDiskMode(b *testing.B) {
	db := state.MakeBadgerDb("scope1", fmt.Sprintf("%s/%s", b.TempDir(), fmt.Sprintf("batch%d", b.N)))
	benchmarkRunner(b, db)
}

func benchmarkRunner(b *testing.B, db state.Db) {
	query := `
	 event e0
then event e1
	 group by session 
`
	ses := parser.ParseSESQuery(query)
	r := MakeRunner(ses, db)

	for i := 0; i < b.N; i++ {
		e := SimpleEvent{
			AnyData{
				"id":      fmt.Sprintf("%d", i),
				"name":    fmt.Sprintf("e%d", i%2), // e0 or e1
				"time":    time.Now().Nanosecond(),
				"session": "fixed",
			},
		}
		r.Accept(&e)
	}

}
