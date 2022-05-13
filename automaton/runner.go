package automaton

import (
	"ses_pm_antlr/automaton/state"
	"ses_pm_antlr/ses"
)

// Runner should accept incoming event and run all available instances of automatons against it,
// if nondeterminism found it should add forked instances to the loop
type Runner struct {
	db        state.Db
	ses       *ses.SES
	instances []*Automaton
}

// removeInstanceAt is internal function to remove an instance that is in a failed state
// it should try to avoid extra allocations and release unused resources
func (r *Runner) removeInstanceAt(i int) {
	r.instances[i] = nil // free up resources
	r.instances = append(r.instances[:i], r.instances[i+1:]...)
}

// removeInstance is internal function to remove an instance that is in a failed state
func (r *Runner) removeInstance(a1 *Automaton) {
	for i, a2 := range r.instances {
		if a2 == a1 {
			r.removeInstanceAt(i)
		}
	}
}

// appendInstance is internal function to add a new instance to the loop (as a result of nondeterminism)
// it should try to avoid extra allocations and release unused resources
func (r *Runner) appendInstance(a *Automaton) {
	r.instances = append(r.instances, a) // no check for duplication
}

// Accept pushes the event to all available automaton instances
func (r *Runner) Accept(e Event) {
	// Every event starts a new potential thread
	startInstance := MakeAutomaton(r.ses, r.db)
	r.appendInstance(startInstance)

	i := 0
	for {
		if i == len(r.instances) {
			break // oob
		}
		_, isFailed, forks := r.instances[i].accept(e)
		for _, f := range forks {
			r.appendInstance(f)
		}
		if isFailed {
			r.removeInstanceAt(i)
			continue
		}
		i = i + 1
	}

	// startInstance is removed if it did not accept the event
	if len(startInstance.GetAcceptedEventIdsAsSlice()) == 0 {
		r.removeInstance(startInstance)
	}
}

// GetAcceptingAutomatons returns automatons that are in accepting state
func (r *Runner) GetAcceptingAutomatons() []*Automaton {
	accepted := make([]*Automaton, 0)
	for _, a := range r.instances {
		if a.IsAcceptingState() {
			accepted = append(accepted, a)
		}
	}
	return accepted
}

func MakeRunner(ses *ses.SES, db state.Db) *Runner {
	return &Runner{
		ses:       ses,
		instances: make([]*Automaton, 0),
		db:        db,
	}
}
