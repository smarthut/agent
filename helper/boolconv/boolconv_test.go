package boolconv_test

import (
	"testing"

	"github.com/smarthut/agent/helper/boolconv"
)

func TestBtoi(t *testing.T) {
	tt := []struct {
		name string
		args bool
		want int
	}{
		{"test 1", true, 1},
		{"test 2", false, 0},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := boolconv.Btoi(tc.args); got != tc.want {
				t.Errorf("Btoi() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestItob(t *testing.T) {
	tt := []struct {
		name string
		args int
		want bool
	}{
		{"test 1", 1, true},
		{"test 2", 0, false},
		{"test 3", -1, true},
		{"test 4", 2, true},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := boolconv.Itob(tc.args); got != tc.want {
				t.Errorf("Itob() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestGetBool(t *testing.T) {
	tt := []struct {
		name string
		args interface{}
		want bool
	}{
		{"test 1", 1, true},
		{"test 2", 0, false},
		{"test 3", true, true},
		{"test 4", false, false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if gotResult, err := boolconv.GetBool(tc.args); gotResult != tc.want || err != nil {
				t.Errorf("GetBool() = %v, %v, want %v", gotResult, err, tc.want)
			}
		})
	}
}
