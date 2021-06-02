package src

import (
	"reflect"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"[1,2,3]", []int{1, 2, 3}, []int{1, 2, 3}},
		{"[1,0,3]", []int{1, 0, 3}, []int{1, 3, 0}},
		{"[0,1,0,3,12]", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
