package src

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{"[1,2,3,4,5]", 11, []int{1, 2, 3, 4, 5}, 3},
		{"[2,3,1,2,4,3]", 7, []int{2, 3, 1, 2, 4, 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.target, tt.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
