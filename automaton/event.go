package automaton

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Event is something that is being analyzed by the automaton
// It can have various attributes that is tested against the query
type Event interface {
	Id() string           // unique id of the event
	Name() string         // a name of the event schema
	Time() time.Time      // the date of occurrence, ordering factor
	Attribute(string) any // get the attribute of the event, a path can have dots "names.first"
}

// SimpleEvent implementation for tests and simple cases
type SimpleEvent struct {
	Data AnyData
}

func (s *SimpleEvent) Id() string {
	idString, isString := s.Data.Read("id").(string)
	if !isString {
		panic(fmt.Sprintf("id not found in event %v", s.Data))
	}
	return idString
}
func (s *SimpleEvent) Name() string {
	idString, isString := s.Data.Read("name").(string)
	if !isString {
		panic(fmt.Sprintf("name not found in event %v", s.Data))
	}
	return idString
}
func (s *SimpleEvent) Time() time.Time {
	v := s.Data.Read("time")
	var t64 int64
	// Assume nanoseconds
	// Convert the input to in64 nanoseconds
	switch t := v.(type) {
	case int:
		t64 = int64(t)
	case int64:
		t64 = int64(t)
	case float64:
		t64 = int64(t)
	case string:
		u, err := strconv.ParseInt(t, 10, 64)
		if err != nil {
			panic(fmt.Sprintf("time has invalid format in event %v, must be unixtime", s.Data))
		}
		t64 = int64(u)
	default:
		panic(fmt.Sprintf("time not found or in invalid format in event %v", s.Data))
	}
	// convert the int64 time to time
	return time.Unix(int64(t64)/int64(time.Second), int64(t64)%int64(time.Second))
}
func (s *SimpleEvent) Attribute(path string) any {
	return s.Data.Read(path)
}

func MakeSimpleEventFromJson(text string) SimpleEvent {
	var d AnyData
	err := json.Unmarshal([]byte(text), &d)
	if err != nil {
		panic(fmt.Sprintf("json decoding failed: %v", err))
	}
	return SimpleEvent{d}
}

type AnyData map[string]any

// Read Data by path (dot-notation)
func (d AnyData) Read(path string) any {
	keys := strings.Split(path, ".")
	return d.read(keys)
}

func (d AnyData) read(nestedKeys []string) any {
	first := nestedKeys[0]
	nestedKeys = nestedKeys[1:]

	val, exists := d[first]
	if !exists {
		return nil
	}

	switch val := val.(type) {
	case AnyData:
		return val.read(nestedKeys)
	case map[string]any:
		return AnyData(val).read(nestedKeys)
	case []any:
		// if there are 2 more keys left - fail
		if len(nestedKeys) > 1 {
			return nil // unable to address nested slices
		}
		if len(nestedKeys) == 1 {
			// the key must be integer, otherwise fail
			index, err := strconv.Atoi(nestedKeys[0])
			if err != nil {
				return nil // the key is not integer, fail
			}
			if index < 0 || index > len(val)-1 {
				return nil // oob fail
			}
		}
		return val // return slice as the requested value
	default:
		if len(nestedKeys) > 1 {
			return nil // unable to address nested scalars
		}
		return val
	}
}
