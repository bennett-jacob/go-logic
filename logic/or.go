package logic

// Or takes a variable number of boolean statements / expressions and returns whether or not one of them was true. This is the logical negation of Nor.
//
// If input is empty, this returns false. To learn why, see "Vacuous Truth": https://en.wikipedia.org/wiki/Vacuous_truth
func Or(propositions ...bool) bool {
	for _, proposition := range propositions {
		if proposition {
			return true
		}
	}
	return false
}

// Any takes a variable number of boolean statements / expressions and returns whether or not one of them was true. This is the logical negation of None.
//
// This is equivalent to Or().
//
// If input is empty, this returns false. To learn why, see "Vacuous Truth": https://en.wikipedia.org/wiki/Vacuous_truth
var Any = Or
