package logic

// None takes a variable number of boolean statements / expressions and returns whether or not every one of them was false, thus all false together.
//
// If input is empty, this returns true, as it is the logical negation of Any.
func None(propositions ...bool) bool {
	return !Any(propositions...)
}
