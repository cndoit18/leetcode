package src

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		wantNums []int
		want     int
	}{
		{"[0,1,2,2,3,0,4,2]", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 4, 0, 3}, 5},
		{"[3,2,2,3]", []int{3, 2, 2, 3}, 3, []int{2, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.nums, tt.val); got != tt.want && reflect.DeepEqual(tt.nums[:got], tt.wantNums) {
				t.Errorf("removeElement() = %v, want %v, nums = %v, want nums %v", got, tt.want, tt.nums[:got], tt.wantNums)

			}
		})
	}
}
