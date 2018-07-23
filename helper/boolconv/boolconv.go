package boolconv

// Btoi converts value bool to int
func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// Itob converts value int value to bool
func Itob(i int) bool {
	return i != 0
}
