package boolconv

import "testing"

func TestBtoi(t *testing.T) {
	tests := []struct {
		name string
		args bool
		want int
	}{
		// TODO: Add test cases.
		{"test 1", true, 1},
		{"test 2", false, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Btoi(test.args); got != test.want {
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
			if got := Itob(test.args); got != test.want {
				t.Errorf("Itob() = %v, want %v", got, test.want)
			}
		})
	}
}
