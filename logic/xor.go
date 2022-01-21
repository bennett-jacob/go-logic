package logic

// OnlyOne takes a variable number of boolean statements / expressions and returns whether or not exactly one of them was true.
func OnlyOne(propositions ...bool) bool {
	numTrue := 0
	for _, proposition := range propositions {
		if proposition {
			numTrue++
		}
	}
	return numTrue == 1
}

// Xor takes a variable number of boolean statements / expressions and returns whether or not exactly one of them was true.
//
// This is equivalent to OnlyOne().
var Xor = OnlyOne
