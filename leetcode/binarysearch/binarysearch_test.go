package binarysearch

import "testing"

func TestSearch(t *testing.T) {

	nums := []int{1, 3, 4, 6, 7, 8, 9}
	t.Log(search(nums, 4))

}
