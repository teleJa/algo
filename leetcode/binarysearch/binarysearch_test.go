package binarysearch

import "testing"

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
