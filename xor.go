package logic

// Xor returns the exclusive or of a and b.
func Xor(a, b bool) bool {
	if a != b {
		return true
	}
	return false
}
