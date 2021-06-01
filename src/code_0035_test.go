package src

import "testing"

func Test_searchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"[-1,0,3,5,9,12]", []int{-1, 0, 3, 5, 9, 12}, 1, 2},
		{"[-1,0,3,5,9,12]", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"[-1,0,3,5,9,12]", []int{-1, 0, 3, 5, 9, 12}, -2, 0},
		{"[1, 3, 5, 6]", []int{1, 3, 5, 6}, 5, 2},
		{"[1,3,5,6]", []int{1, 3, 5, 6}, 2, 1},
		{"[1,3,5,6]", []int{1, 3, 5, 6}, 7, 4},
		{"[1,3,5,6]", []int{1, 3, 5, 6}, 0, 0},
		{"[1,3]", []int{1, 3}, 2, 1},
		{"[1,3]", []int{1, 3}, 4, 2},
		{"[1,3]", []int{1, 3}, 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.nums, tt.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}
