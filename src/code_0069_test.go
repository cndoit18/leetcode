package src

import "testing"

func Test_mySqrt(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"0", 0, 0},
		{"4", 4, 2},
		{"6", 6, 2},
		{"8", 8, 2},
		{"9", 9, 3},
		{"10", 10, 3},
		{"15", 15, 3},
		{"18", 18, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
