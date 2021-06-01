package src

import (
	"reflect"
	"testing"
)

func Test_searchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"[5,7,7,8,8,10]", []int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{"[5,7,7,8,8,10]", []int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{"[]", []int{}, 0, []int{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
