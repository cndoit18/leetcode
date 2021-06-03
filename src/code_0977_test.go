package src

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"[-4,-1,0,3,10]", []int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{"[-1,1,2,4]", []int{-1, 1, 2, 4}, []int{1, 1, 4, 16}},
		{"[-4,1,2,4]", []int{-4, 1, 2, 4}, []int{1, 4, 16, 16}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
