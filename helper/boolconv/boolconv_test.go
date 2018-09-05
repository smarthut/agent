package boolconv_test

import (
	"testing"

	"github.com/smarthut/agent/helper/boolconv"
)

func TestBtoi(t *testing.T) {
	tests := []struct {
		name string
		args bool
		want int
	}{
		{"test 1", true, 1},
		{"test 2", false, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := boolconv.Btoi(test.args); got != test.want {
				t.Errorf("Btoi() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestItob(t *testing.T) {
	tests := []struct {
		name string
		args int
		want bool
	}{
		{"test 1", 1, true},
		{"test 2", 0, false},
		{"test 3", -1, true},
		{"test 4", 2, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := boolconv.Itob(test.args); got != test.want {
				t.Errorf("Itob() = %v, want %v", got, test.want)
			}
		})
	}
}

func TestGetBool(t *testing.T) {
	type want struct {
		result bool
		err    error
	}
	tests := []struct {
		name string
		args interface{}
		want
	}{
		{"test 1", 1, want{true, nil}},
		{"test 2", 0, want{false, nil}},
		{"test 3", true, want{true, nil}},
		{"test 4", false, want{false, nil}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if gotResult, gotErr := boolconv.GetBool(test.args); gotResult != test.want.result || gotErr != test.want.err {
				t.Errorf("GetBool() = %v, %v, want %v", gotResult, gotErr, test.want)
			}
		})
	}
}
