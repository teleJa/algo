package leetcode

import (
	"log"
	"testing"
)

func TestMoveZeros(t *testing.T) {

	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	log.Println(arr)
}
