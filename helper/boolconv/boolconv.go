package boolconv

import (
	"fmt"
)

// Errors
var (
	ErrUnableToCastValue = "boolconv: unable to cast %v to bool value"
)

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

// GetBool casts interface to a bool value
func GetBool(v interface{}) (bool, error) {
	switch v.(type) {
	case bool:
		return v.(bool), nil
	case int:
		return Itob(v.(int)), nil
	}
	return false, fmt.Errorf(ErrUnableToCastValue, v)
}
