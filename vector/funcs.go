package vector

func castFloat64(eventAttribute any) (float64, bool) {
	eventAttributeFloat, isFloat := eventAttribute.(float64)
	if !isFloat {
		return .0, false // the attribute can not be compared with the float
	}
	return eventAttributeFloat, true
}

// anyMatch returns true if at least one match is found
func anyMatch(value any, vectorIterator Iterator, cmpFunc func(a, b any) bool) bool {
	for {
		vectorValue := vectorIterator()
		if vectorValue == nil {
			break // iteration is over
		}

		if cmpFunc(value, vectorValue) {
			return true
		}
	}
	return false
}

func anyMatchFloat64(value float64, vectorIterator Iterator, cmpFunc func(a, b float64) bool) bool {
	for {
		vectorValue := vectorIterator()
		_, vectorValueIsFloat := vectorValue.(float64) // dynamic type check
		if vectorValue == nil || !vectorValueIsFloat {
			break // iteration is over
		}

		vectorValueFloat64, ok := vectorValue.(float64) // todo this type check can be eliminated
		if !ok {
			continue // try the next one
		}

		if cmpFunc(value, vectorValueFloat64) {
			return true
		}
	}
	return false
}

// allMatch returns true if all vector values match
func allMatch(value any, vectorIterator Iterator, cmpFunc func(a, b any) bool) bool {
	for {
		vectorValue := vectorIterator()
		_, vectorValueIsFloat := vectorValue.(float64) // dynamic type check
		if vectorValue == nil || !vectorValueIsFloat {
			break // iteration is over
		}
		if !cmpFunc(value, vectorValue) {
			return false
		}
	}
	return true
}

func allMatchFloat64(value float64, vectorIterator Iterator, cmpFunc func(a, b float64) bool) bool {
	for {
		vectorValue := vectorIterator()
		if vectorValue == nil {
			break // iteration is over
		}

		vectorValueFloat64, ok := vectorValue.(float64)
		if !ok {
			return false // one of the values is not a float, so return false
		}

		if !cmpFunc(value, vectorValueFloat64) {
			return false
		}
	}
	return true
}
