package automaton

import (
	"fmt"
	"reflect"
	"testing"
)

func TestJsonParse(t *testing.T) {
	type testInput struct {
		json          string
		path          string
		expectedValue any
	}

	tests := []testInput{
		{`{}`, "name", nil},
		{`{}`, "name.first", nil},
		{`{"name": "Dmitriy"}`, "name", "Dmitriy"},
		{`{"number": 100}`, "number", 100.00},
		{`{"period": {"range": true}}`, "period.range", true},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			se := MakeSimpleEventFromJson(tt.json)
			actual := se.Attribute(tt.path)
			if !reflect.DeepEqual(actual, tt.expectedValue) {
				t.Errorf("Test %d failed: actual %v(%[2]T) and expected %v(%[3]T)", i, actual, tt.expectedValue)
			}
		})
	}
}

func TestAnyDataRead(t *testing.T) {
	testData := []struct {
		anydata       AnyData
		path          string
		expectedValue any
	}{
		{AnyData{}, "name", nil},
		{AnyData{}, "name.first", nil},
		{AnyData{"name": "Dmitriy"}, "name", "Dmitriy"},
		{AnyData{"number": 100}, "number", 100},
		{AnyData{"period": AnyData{"range": true}}, "period.range", true},
	}

	for i, tt := range testData {
		actual := tt.anydata.Read(tt.path)
		if !reflect.DeepEqual(actual, tt.expectedValue) {
			t.Errorf("Test %d failed: %v(%[2]T) and %v(%[3]T)", i, actual, tt.expectedValue)
		}
	}
}
