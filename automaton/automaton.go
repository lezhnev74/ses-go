package automaton

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

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

	// spec:
	ses         *ses.SES                               // source definition of the automaton
	matchers    map[int]map[string]expr.BoolExpression // map matches to set events
	bufferSpecs map[string]map[string]attrBufferSpec   // a map from event to attr name to buffer spec, when we need to capture an attr we use this spec to make a buffer

	// state:
	id                 uint64                             // unique id
	failed             bool                               // set when nothing can no longer be accepted
	curSet             int                                // current set index
	capturedAttributes map[int]map[string]EventAttrBuffer // map full attr name to buffer
	acceptedEventIds   map[int]map[string]EventAttrBuffer // ids per set and per event
	setLastEventTime   map[int]*time.Time                 // remember the last event's time for each set (used in window matching)
	groupBy            []any                              // keep the unique value that we use to group events (populated from the initial event)
	windowStart        *time.Time                         // first captured event's time
}

// accept tests the incoming event according to automaton criteria
// returns isAccepted if event was accepted otherwise false, isFailed is true if automaton reached unrecoverable fail state, forks are to test alternative branches of matching
// note: if isFailed is true then isAccepted is false
func (a *Automaton) accept(incomingEvent Event) (isAccepted bool, isFailed bool, forks []*Automaton) {
	// Condition: If automaton is in failed state - reject
	if a.failed {
		return
	}

	// Condition: events must share group-by attributes
	if !a.checkGroupBy(incomingEvent) {
		return
	}

	// Condition: event must remain within the SES window
	windowMatch := a.checkSesWindow(incomingEvent)
	if windowMatch != 0 {

		// Edge case: window is over
		if windowMatch > 0 {
			if !a.IsAcceptingState() {
				// Assume events are ordered and no following events can fit in the window
				// So if the current automaton is NOT accepting then it won't ever be
				a.failed = true
				isFailed = true
			}
		}

		isAccepted = false
		return
	}

	// NonDeterminism Condition:
	// 3.1 if current set is in ACCEPTING state, check if event matches the next set, if yes - create forked automaton
	if a.isAcceptingSet(a.curSet) && a.matchEventInSet(incomingEvent, a.curSet) && a.matchEventInSet(incomingEvent, a.curSet+1) {
		forks = append(forks, a.fork())
	}

	// Condition: Apply event condition expression, if fails - check the next set otherwise reject
	if !a.matchEventInSet(incomingEvent, a.curSet) {
		// if the current set is accepting but event does not match it, then try the next set in the automaton
		if a.isAcceptingSet(a.curSet) && a.matchEventInSet(incomingEvent, a.curSet+1) {
			a.curSet = a.curSet + 1
		} else {
			return
		}
	}

	// Put event to buffers (id, attrs)
	isAccepted = true
	a.saveAcceptedEvent(incomingEvent)

	// Condition: Apply qty check, if underflow - accept, wait for more events, if overflow - put the automaton to FAILED state and reject
	for _, event := range a.ses.GetSets()[a.curSet].GetEvents() {
		if _, isOverflow := a.checkEventQty(a.curSet, event); isOverflow {
			a.failed = true
			isAccepted = false
			isFailed = true
			return
		}
	}

	return
}

// checkSesWindow return -1 is event is below the window, 0 if withing the window, 1 if after the window
// the func has side-effects (it changes windowStart attribute)
func (a *Automaton) checkSesWindow(e Event) int {
	eventTime := e.Time()

	// Condition: skip window check if duration is 0
	windowDuration := a.ses.GetWindow().GetDuration()
	if windowDuration == 0 {
		return 0
	}

	// Condition: if this automaton is fresh and has no window initialized yet
	if a.windowStart == nil {
		if a.ses.GetWindow().IsSliding() {
			a.windowStart = &eventTime
		} else { // absolute bounds
			from, _ := a.ses.GetWindow().GetBounds()
			a.windowStart = &from
		}
		return 0
	}

	// Condition: if an event is before the window (wrong event order...)
	if eventTime.Before(*a.windowStart) {
		return -1
	}

	// Condition: if an event after the window
	windowEnd := a.windowStart.Add(windowDuration)
	if eventTime.After(windowEnd) {
		return 1
	}

	return 0 // otherwise all is good
}

// checkSetWindow returns true if the current event fits into the set window
// only works from the 2nd set on
func (a *Automaton) checkSetWindow(e Event, set int) bool {
	// Edge case: the first set does not need this check
	if set == 0 {
		return true
	}

	lastEventTime, captured := a.setLastEventTime[set-1]
	// Edge case: no events captured in the previous set, so window checking is not applicable
	// THis can happen in a query like "event a? then within 1 minute event b"
	if !captured {
		return true
	}

	window := a.ses.GetSets()[set].GetWindow()
	eventTime := e.Time()
	windowStart := lastEventTime.Add(window.Skip)
	windowEnd := windowStart.Add(window.Within)

	// Condition: SKIP criteria
	if eventTime.Before(windowStart) {
		return false // the event is too early
	}

	if window.Within != 0 && eventTime.After(windowEnd) {
		return false // the event is too late
	}

	return true
}

// Group By are attributes that must have the same value for every event
func (a *Automaton) checkGroupBy(e Event) bool {
	if len(a.groupBy) == 0 { // first event, must accept
		return true
	}

	actualValues := make([]any, 0, len(a.ses.GetGroupBy()))
	for _, attr := range a.ses.GetGroupBy() {
		actualValues = append(actualValues, e.Attribute(attr))
	}
	return vector.CmpAnySlices(actualValues, a.groupBy)
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

// saveAcceptedEvent saves matched event to automaton buffers
func (a *Automaton) saveAcceptedEvent(e Event) {
	// save group by
	if len(a.groupBy) == 0 {
		actualValues := make([]any, 0, len(a.ses.GetGroupBy()))
		for _, attr := range a.ses.GetGroupBy() {
			actualValues = append(actualValues, e.Attribute(attr))
		}
		a.groupBy = actualValues
	}

	t := e.Time()
	a.setLastEventTime[a.curSet] = &t
	eventName := a.resolveEventName(e)

	// save capturing attributes
	for event, attrSpecs := range a.bufferSpecs {
		if event != eventName {
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

	_, exists = a.acceptedEventIds[a.curSet][eventName]
	if !exists {
		a.acceptedEventIds[a.curSet][eventName] = MakeSimpleEventAttrBuffer(a.db)
	}
	a.acceptedEventIds[a.curSet][eventName].Accept(e.Id())
}

// matchEventInSet tests a given event against criteria of a given set
// criteria includes window-check and where-clauses check
func (a *Automaton) matchEventInSet(e Event, set int) bool {
	// Check: set events exposes matching functions, if no - then event is not in the set
	setConds, exists := a.matchers[set]
	if !exists {
		return false
	}
	eventName := a.resolveEventName(e)
	eventCond, exists := setConds[eventName]
	if !exists {
		return false
	}

	// Check: run where-criteria
	env := &AutomatonEnv{a, e}
	if !eventCond(env) {
		return false
	}

	// Check window criteria starting from the second set
	if !a.checkSetWindow(e, set) {
		return false
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

func (a *Automaton) GetGroupBy() any {
	return a.groupBy
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

// resolveEventName returns name of the incoming event with respect to aliasing
func (a *Automaton) resolveEventName(e Event) string {
	for _, setEvent := range a.ses.GetSets()[a.curSet].GetEvents() {
		if setEvent.GetOriginalName() == e.Name() {
			return setEvent.GetName()
		}
	}
	return e.Name()
}

// save serializes the current state of the automaton and saves it to the db
func (a *Automaton) save() {
	serialized := make(map[string]any) // put all the data in here
	serialized["id"] = a.id
	serialized["groupBy"] = a.groupBy
	serialized["failed"] = a.failed
	serialized["curSet"] = a.curSet

	serialized["windowStart"] = nil
	if a.windowStart != nil {
		tt, err := a.windowStart.MarshalText()
		if err != nil {
			panic(err)
		}
		serialized["windowStart"] = tt
	}

	setTimes := make(map[int]string)
	for s, t := range a.setLastEventTime {
		if t != nil {
			tt, err := t.MarshalText()
			if err != nil {
				panic(err)
			}
			setTimes[s] = string(tt)
		}
	}
	serialized["setLastEventTime"] = setTimes

	capturedAttrs := make(map[int]map[string]string)
	for s, eventBuf := range a.capturedAttributes {
		capturedAttrs[s] = make(map[string]string)
		for event, attrBuf := range eventBuf {
			capturedAttrs[s][event] = attrBuf.Serialize()
		}
	}
	serialized["capturedAttrs"] = capturedAttrs

	acceptedEventIds := make(map[int]map[string]string)
	for s, eventBuf := range a.acceptedEventIds {
		acceptedEventIds[s] = make(map[string]string)
		for event, attrBuf := range eventBuf {
			acceptedEventIds[s][event] = attrBuf.Serialize()
		}
	}
	serialized["acceptedEventIds"] = acceptedEventIds

	automatonState, err := json.Marshal(serialized)
	if err != nil {
		panic(err)
	}

	a.db.Save(a.id, automatonState)
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

func MakeAutomaton(s *ses.SES, db state.Db) *Automaton {
	a := Automaton{
		id:               db.NextId(),
		ses:              s,
		matchers:         make(map[int]map[string]expr.BoolExpression),
		db:               db,
		setLastEventTime: make(map[int]*time.Time),
	}

	// 1. Make buffers
	a.capturedAttributes = make(map[int]map[string]EventAttrBuffer)
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
		a.matchers[i] = make(map[string]expr.BoolExpression)

		for _, event := range eventSets.GetEvents() {
			if event.GetCondition() == nil {
				a.matchers[i][event.GetName()] = func(e expr.Environment) bool { return true } // default always-true
				continue
			}
			// one unified expression for all event WHERE conditions
			a.matchers[i][event.GetName()] = expr.MakeAndBoolExpression(expr.MakeExpressionFromCondition(event.GetCondition()))
		}
	}

	// 3. Finish creation
	return &a
}

func RestoreAutomaton(id uint64, ses *ses.SES, db state.Db) (restoredAutomaton *Automaton) {
	var data map[string]any
	serializedAutomaton := db.Find(id)
	err := json.Unmarshal(serializedAutomaton, &data)
	if err != nil {
		panic(err)
	}

	// Now make a new automaton
	restoredAutomaton = MakeAutomaton(ses, db)
	restoredAutomaton.id = uint64(data["id"].(float64))
	if data["groupBy"] != nil {
		restoredAutomaton.groupBy = data["groupBy"].([]any)
	}
	restoredAutomaton.failed = data["failed"].(bool)
	restoredAutomaton.curSet = int(data["curSet"].(float64))

	if data["windowStart"] != nil {
		n := time.Now()
		err := n.UnmarshalText(data["windowStart"].([]byte))
		if err != nil {
			panic(err)
		}
		restoredAutomaton.windowStart = &n
	}
	if data["setLastEventTime"] != nil {
		m := data["setLastEventTime"].(map[string]any)
		for s, t := range m {
			i, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				panic(err)
			}
			n := time.Now()
			err = n.UnmarshalText([]byte(t.(string)))
			if err != nil {
				panic(err)
			}
			restoredAutomaton.setLastEventTime[int(i)] = &n
		}
	}
	capturedAttrs := data["capturedAttrs"].(map[string]any)
	for s, p1 := range capturedAttrs {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}

		if restoredAutomaton.capturedAttributes[int(i)] == nil {
			restoredAutomaton.capturedAttributes[int(i)] = make(map[string]EventAttrBuffer)
		}

		evs := p1.(map[string]any)
		for ev, bufState := range evs {
			if restoredAutomaton.capturedAttributes[int(i)][ev] == nil {
				subs := strings.SplitN(ev, ".", 2)
				eventName := subs[0]
				eventAttribute := subs[1]
				restoredAutomaton.capturedAttributes[int(i)][ev] = MakeBufferFromSpec(restoredAutomaton.bufferSpecs[eventName][eventAttribute], db)
			}

			restoredAutomaton.capturedAttributes[int(i)][ev].Unserialize(bufState.(string))
		}
	}

	restoredAutomaton.acceptedEventIds = make(map[int]map[string]EventAttrBuffer)
	acceptedEventIds := data["acceptedEventIds"].(map[string]any)
	for s, p1 := range acceptedEventIds {
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		if _, exists := restoredAutomaton.acceptedEventIds[int(i)]; !exists {
			restoredAutomaton.acceptedEventIds[int(i)] = make(map[string]EventAttrBuffer)
		}

		evs := p1.(map[string]any)
		for ev, bufState := range evs {
			if _, exists := restoredAutomaton.acceptedEventIds[int(i)][ev]; !exists {
				restoredAutomaton.acceptedEventIds[int(i)][ev] = MakeSimpleEventAttrBuffer(db)
			}
			restoredAutomaton.acceptedEventIds[int(i)][ev].Unserialize(bufState.(string))
		}
	}

	return
}
