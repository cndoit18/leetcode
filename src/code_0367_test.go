package src

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"16", 16, true},
		{"14", 14, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
