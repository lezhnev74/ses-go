package automaton

import (
	"ses_pm_antlr/automaton/state"
	"ses_pm_antlr/expr"
	"ses_pm_antlr/ses"
	"ses_pm_antlr/vector"
)

// Automaton accepts incoming events and captures matching ones
// it tests the event based on the given SES specification
// It can detect non-determinism and fork to test all possible branches
type Automaton struct {
	db state.Db

	ses                *ses.SES                                          // source definition of the automaton
	matchers           map[int]map[string]map[string]expr.BoolExpression // map matches to set events
	bufferSpecs        map[string]map[string]attrBufferSpec              // a map from event to attr name to buffer spec, when we need to capture an attr we use this spec to make a buffer
	failed             bool                                              // set when nothing can no longer be accepted
	curSet             int                                               // current set index
	capturedAttributes map[int]map[string]EventAttrBuffer                // map full attr name to buffer
	acceptedEventIds   map[int]map[string]EventAttrBuffer                // ids per set and per event
	groupBy            any                                               // keep the unique value that we use to group events
}

// Accept tests the incoming event according to automaton criteria
// returns isAccepted if event was accepted otherwise false, isFailed is true if automaton reached unrecoverable fail state, forks are to test alternative branches of matching
// note: if isFailed is true then isAccepted is false
func (a *Automaton) Accept(incomingEvent Event) (isAccepted bool, isFailed bool, forks []*Automaton) {
	// 1. If automaton is in failed state - reject
	if a.failed {
		return
	}

	// 2. If at least one event captured before, check group-by equality, if fails - reject
	if a.groupBy != nil && incomingEvent.Attribute(a.ses.GetGroupBy()) != a.groupBy {
		return
	}

	// 3. NonDeterminism check
	// 3.1 if current set is in ACCEPTING state, check if event matches the next set, if yes - create forked automaton
	if a.isAcceptingSet(a.curSet) && a.matchEventInSet(incomingEvent, a.curSet) && a.matchEventInSet(incomingEvent, a.curSet+1) {
		forks = append(forks, a.fork())
	}

	// 4. Apply event condition expression, if fails - check the next set otherwise reject
	if !a.matchEventInSet(incomingEvent, a.curSet) {
		// if the current set is accepting but event does not match it, then try the next set in the automaton
		if a.isAcceptingSet(a.curSet) && a.matchEventInSet(incomingEvent, a.curSet+1) {
			a.curSet = a.curSet + 1
		} else {
			return
		}
	}

	// 5. Put event to buffers (id, attrs)
	isAccepted = true
	a.saveAcceptedEvent(incomingEvent)

	// 6. Apply qty check, if underflow - accept, wait for more events, if overflow - put the automaton to FAILED state and reject
	for _, event := range a.ses.GetSets()[a.curSet].GetEvents() {
		if _, isOverflow := a.checkEventQty(a.curSet, event); isOverflow {
			a.failed = true
			isAccepted = false
			isFailed = true
		}
	}

	return
}

// saveAcceptedEvent saves matched event to automaton buffers
func (a *Automaton) saveAcceptedEvent(e Event) {
	a.groupBy = e.Attribute(a.ses.GetGroupBy())

	// save capturing attributes
	for event, attrSpecs := range a.bufferSpecs {
		if event != e.Name() {
			continue
		}

		for _, spec := range attrSpecs {
			val := e.Attribute(spec.operand.EventAttribute)
			if val == nil {
				continue // do not save nils
			}

			// save value to the buffer
			// init buffers if not already
			_, exists := a.capturedAttributes[a.curSet]
			if !exists {
				a.capturedAttributes[a.curSet] = make(map[string]EventAttrBuffer)
			}
			_, exists = a.capturedAttributes[a.curSet][spec.operand.GetQualifiedAttributeName()]
			if !exists {
				a.capturedAttributes[a.curSet][spec.operand.GetQualifiedAttributeName()] = MakeBufferFromSpec(spec, a.db)
			}
			a.capturedAttributes[a.curSet][spec.operand.GetQualifiedAttributeName()].Accept(val)
		}
	}

	// save event ids
	// save value to the buffer
	// init buffers if not already
	_, exists := a.acceptedEventIds[a.curSet]
	if !exists {
		a.acceptedEventIds[a.curSet] = make(map[string]EventAttrBuffer)
	}

	_, exists = a.acceptedEventIds[a.curSet][e.Name()]
	if !exists {
		a.acceptedEventIds[a.curSet][e.Name()] = MakeSimpleEventAttrBuffer(a.db)
	}
	a.acceptedEventIds[a.curSet][e.Name()].Accept(e.Id())
}

// matchEventInSet tests a given event against criteria of a given set
func (a *Automaton) matchEventInSet(e Event, set int) bool {
	setConds, exists := a.matchers[set]
	if !exists {
		return false
	}
	eventConds, exists := setConds[e.Name()]
	if !exists {
		return false
	}

	env := &AutomatonEnv{a, e}
	for _, attrCond := range eventConds {
		if !attrCond(env) {
			return false
		}
	}
	return true
}

// isAcceptingSet checks if captured events for the set satisfy qty criteria
func (a *Automaton) isAcceptingSet(set int) bool {
	for _, event := range a.ses.GetSets()[set].GetEvents() {
		if isMatch, _ := a.checkEventQty(set, event); !isMatch {
			return false
		}
	}
	return true
}

// IsAcceptingState returns true if the whole automaton is in accepting state
func (a *Automaton) IsAcceptingState() bool {
	if a.failed {
		return false
	}
	for i := range a.ses.GetSets() {
		if !a.isAcceptingSet(i) {
			return false
		}
	}
	return true
}

func (a *Automaton) GetAcceptedEventIds() map[int]map[string]EventAttrBuffer {
	return a.acceptedEventIds
}

func (a *Automaton) GetAcceptedEventIdsAsSlice() []map[string][]any {
	ids := make([]map[string][]any, len(a.GetAcceptedEventIds()))
	for set, eventBuffers := range a.GetAcceptedEventIds() {
		ids[set] = make(map[string][]any)
		for eventName, buf := range eventBuffers {
			ids[set][eventName] = buf.GetIterator().ToInvertedSlice() // iteration goes backwards
		}
	}
	return ids
}

// checkEventQty tests the number of captured events against qty criteria in the SES specification
func (a *Automaton) checkEventQty(set int, event *ses.Event) (isMatch, isOverflow bool) {
	min, max := event.GetQty()
	actualCount := 0

	capturedSetIds, exists := a.acceptedEventIds[set]
	if exists {
		capturedEventIds, exists := capturedSetIds[event.GetName()]
		if exists {
			actualCount = len(capturedEventIds.GetIterator().ToSlice())
		}
	}

	if actualCount > max {
		isOverflow = true
	}

	if actualCount >= min && actualCount <= max {
		isMatch = true
	}

	return
}

// fork creates a new instance and switches to the next set
func (a *Automaton) fork() *Automaton {
	fork := *a                    // copy by value
	fork.curSet = fork.curSet + 1 // switch to the next set

	// allocate new memory to buffers
	fork.acceptedEventIds = make(map[int]map[string]EventAttrBuffer)
	fork.capturedAttributes = make(map[int]map[string]EventAttrBuffer)
	// links to previous automaton's buffers
	for set, eventAttrs := range a.acceptedEventIds {
		fork.acceptedEventIds[set] = make(map[string]EventAttrBuffer)
		fork.capturedAttributes[set] = make(map[string]EventAttrBuffer)

		for eventAttr, buf := range eventAttrs {
			fork.acceptedEventIds[set][eventAttr] = buf.Spawn()
			fork.capturedAttributes[set][eventAttr] = buf.Spawn()
		}
	}

	return &fork
}

func MakeAutomaton(s *ses.SES, db state.Db) *Automaton {
	a := Automaton{
		ses:      s,
		matchers: make(map[int]map[string]map[string]expr.BoolExpression),
		db:       db,
	}

	// 1. Make buffers
	a.bufferSpecs = make(map[string]map[string]attrBufferSpec)
	for _, spec := range mapSesToSpecs(s) {
		if _, exists := a.bufferSpecs[spec.operand.EventName]; !exists {
			a.bufferSpecs[spec.operand.EventName] = make(map[string]attrBufferSpec)
		}
		a.bufferSpecs[spec.operand.EventName][spec.operand.EventAttribute] = spec
	}
	a.acceptedEventIds = make(map[int]map[string]EventAttrBuffer)

	// 2. Make condition expressions
	for i, eventSets := range s.GetSets() {
		a.matchers[i] = make(map[string]map[string]expr.BoolExpression)

		for _, event := range eventSets.GetEvents() {
			a.matchers[i][event.GetName()] = make(map[string]expr.BoolExpression)

			eventAttrConds := map[string]map[string][]expr.Expression{} // collect all cond expressions for every attribute
			for _, cond := range event.GetConditions() {
				curAttrOperand := cond.GetLeftOperand().(ses.EventAttributeOperand)

				_, exists := eventAttrConds[curAttrOperand.EventName][curAttrOperand.EventAttribute]
				if !exists {
					eventAttrConds[curAttrOperand.EventName] = make(map[string][]expr.Expression)
				}

				eventAttrConds[curAttrOperand.EventName][curAttrOperand.EventAttribute] = append(eventAttrConds[curAttrOperand.EventName][curAttrOperand.EventAttribute], expr.MakeExpressionFromCondition(cond))
			}

			for eventName, eventExprs := range eventAttrConds {

				for attrName, attrExprs := range eventExprs {
					a.matchers[i][eventName][attrName] = expr.MakeAndBoolExpression(attrExprs...) // join all event attribute conditions with AND
				}
			}
		}
	}

	// 3. Finish creation
	return &a
}

// AutomatonEnv allows to use Automaton as an env for expr evaluation
type AutomatonEnv struct {
	a        *Automaton
	curEvent Event
}

func (env *AutomatonEnv) Resolve(operand ses.EventAttributeOperand) any {
	if operand.IsCurrent { // scalar value
		val := env.curEvent.Attribute(operand.EventAttribute)
		return val
	}

	// captured vector value
	val := env.a.capturedAttributes[env.a.curSet][operand.GetQualifiedAttributeName()]
	if val == nil {
		return vector.MakeLinkedSlice(nil).GetIterator()
	}
	return val.GetIterator()
}

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
		_, isFailed, forks := r.instances[i].Accept(e)
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
