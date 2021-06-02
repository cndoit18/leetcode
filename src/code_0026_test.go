package src

import "testing"

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"[1,1,2]", []int{1, 1, 2}, 2},
		{"[1,2]", []int{1, 2}, 2},
		{"[1,1]", []int{1, 1}, 1},
		{"[1]", []int{1}, 1},
		{"[]", []int{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
