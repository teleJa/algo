package leetcode

import (
	"log"
	"testing"
)

func TestMerge(t *testing.T) {

	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3

	n4 := &ListNode{Val: 1}
	n5 := &ListNode{Val: 3}
	n6 := &ListNode{Val: 4}
	n4.Next = n5
	n5.Next = n6

	lists := mergeTwoLists(n1, n4)
	log.Println(lists)

}
