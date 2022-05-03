package automaton

import (
	"ses_pm_antlr/ses"
)

// attrBufferSpec is a specification for a buffer of captured attributes, it controls the life cycle of a buffer
// setScope determines the scope of the buffer, when the automaton transitions to a set where this buffer is no longer used, the buffer can be removed
type attrBufferSpec struct {
	operand  ses.EventAttributeOperand // parsed operand
	setScope int                       // (set scope) set index in which this buffer is used, any later set does not need this buffer no more
}

// makeBufferSpecs analyzes all conditions in SES and finds attributes that must be captured for matching operations
func makeBufferSpecs(s *ses.SES) []attrBufferSpec {
	// 1. Map all conditions to buffer specs
	allSpecs := mapSesToSpecs(s)

	// 2. Post process: collapse specs for the same attribute:
	attrSpecs := make(map[string]attrBufferSpec, 0)
	for _, spec := range allSpecs {
		existingSpec, exists := attrSpecs[spec.operand.GetQualifiedAttributeName()]
		if !exists {
			attrSpecs[spec.operand.GetQualifiedAttributeName()] = spec
			continue
		}

		// ok here we need to collapse the two conditions for the same attribute
		mergedSpec := existingSpec // make a copy

		// 2.1. the duration must span over the most recent set
		mergedSpec.setScope = selectExpiration(existingSpec, spec)

		// 2.2. the value extraction should be the broadest (all>any, any>first, etc...)
		mergedSpec.operand.SelectMode = selectMode(existingSpec, spec)

		// 3. Place the merged copy
		attrSpecs[spec.operand.GetQualifiedAttributeName()] = mergedSpec
	}

	// 3. Map to slice again
	finalSpecs := make([]attrBufferSpec, 0)
	for _, spec := range attrSpecs {
		finalSpecs = append(finalSpecs, spec)
	}

	return finalSpecs
}

func selectMode(spec1, spec2 attrBufferSpec) int8 {
	// power determines the power of the value extractor set
	power := map[int8]int8{
		ses.SelectAll: 2, // includes every value
		ses.SelectAny: 1, // some
		// anything else is 0 (includes just a single value, like first, last, avg, median etc.)
	}

	// 1. Edge case: if the value extractors are the same, then skip this step altogether
	if spec1.operand.SelectMode == spec2.operand.SelectMode {
		return spec1.operand.SelectMode
	}

	// 2. If the value extractors are different we need to make a choice
	// 2.1 Case: two values of power=0 makes it ANY (so increases the power by 1)
	if power[spec1.operand.SelectMode] == 0 && power[spec2.operand.SelectMode] == 0 {
		return ses.SelectAny
	}

	// 2.2 Otherwise return the biggest power
	if power[spec1.operand.SelectMode] > power[spec2.operand.SelectMode] {
		return spec1.operand.SelectMode
	} else {
		return spec2.operand.SelectMode
	}
}

func selectExpiration(spec1, spec2 attrBufferSpec) int {
	if spec1.setScope > spec2.setScope {
		return spec1.setScope
	}
	return spec2.setScope
}

func mapSesToSpecs(s *ses.SES) []attrBufferSpec {
	specs := make([]attrBufferSpec, 0)
	for setIndex, eventSet := range s.GetSets() {
		for _, event := range eventSet.GetEvents() {
			for _, c := range event.GetConditions() {
				// only test the right operand as the left one is always normalized to be the current event attribute
				rightEventAttribute, rOk := c.GetRightOperand().(ses.EventAttributeOperand)
				if rOk {
					specs = append(specs, attrBufferSpec{rightEventAttribute, setIndex})
				}
			}
		}
	}
	return specs
}
