package binarysearch

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {

	nums := []int{-1, 0, 3, 5, 9, 12}
	t.Log(search(nums, 9))

}

func TestLeftBound(t *testing.T) {

	nums := []int{-1, 0, 3, 5, 7, 9, 12}
	t.Log(leftBound(nums, 8)) // 5

}

func TestRightBound(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 7, 9, 12}
	t.Log(rightBound(nums, 8)) // 4
}

func TestSearchRange(t *testing.T) {

	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			[]int{5, 7, 7, 8, 8, 10},
			8,
			[]int{3, 4},
		},
		{
			[]int{5, 7, 7, 8, 8, 10},
			7,
			[]int{1, 2},
		},
		{
			[]int{},
			0,
			[]int{-1, -1},
		},
	}

	for i, t0 := range tests {
		t.Run(fmt.Sprint("case", i), func(t *testing.T) {

			if !reflect.DeepEqual(searchRange(t0.nums, t0.target), t0.want) {
				t.Error("test fail")
			}

		})
	}

}
