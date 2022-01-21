package logic

// And takes a variable number of boolean statements / expressions and returns whether or not every one of them was true, thus all true together. This is NOT the logical negation of Nor.
//
// If input is empty, this returns true. To learn why, see "Vacuous Truth": https://en.wikipedia.org/wiki/Vacuous_truth
func And(propositions ...bool) bool {
	for _, proposition := range propositions {
		if !proposition {
			return false
		}
	}
	return true
}

// All takes a variable number of boolean statements / expressions and returns whether or not every one of them was true, thus all true together. This is NOT the logical negation of None.
//
// This is equivalent to And().
//
// If input is empty, this returns true. To learn why, see "Vacuous Truth": https://en.wikipedia.org/wiki/Vacuous_truth
var All = And
