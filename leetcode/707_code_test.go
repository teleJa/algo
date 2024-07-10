package leetcode

import (
	"log"
	"testing"
)

func TestFirst(t *testing.T) {

	list := Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	log.Println("index 1:", list.Get(1))
	list.DeleteAtIndex(1)
	log.Println("index 1:", list.Get(1))
}
