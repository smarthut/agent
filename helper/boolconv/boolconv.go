package boolconv

import (
	"fmt"
	"strconv"
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
func GetBool(i interface{}) (bool, error) {
	switch v := i.(type) {
	case bool:
		return v, nil
	case int:
		return Itob(v), nil
	case string:
		return strconv.ParseBool(v)
	}
	return false, fmt.Errorf(ErrUnableToCastValue, i)
}
