package logic

// Nor takes a variable number of boolean statements / expressions and returns whether or not every one of them was false, thus all false together.
//
// If input is empty, this returns true, as it is the logical negation of Or.
func Nor(propositions ...bool) bool {
	return !Any(propositions...)
}

// None takes a variable number of boolean statements / expressions and returns whether or not every one of them was false, thus all false together.
//
// This is equivalent to Nor().
//
// If input is empty, this returns true, as it is the logical negation of Any.
var None = Nor
