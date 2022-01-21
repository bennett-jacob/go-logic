package logic

// Any takes a variable number of boolean statements / expressions and returns whether or not one of them was true. This is the logical negation of None.
//
// If input is empty, this returns false. To learn why, see "Vacuous Truth": https://en.wikipedia.org/wiki/Vacuous_truth
func Any(propositions ...bool) bool {
	for _, proposition := range propositions {
		if proposition {
			return true
		}
	}
	return false
}
