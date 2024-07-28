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

func TestPartition(t *testing.T) {

	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 4}
	n1.Next = n2

	n3 := &ListNode{Val: 3}
	n2.Next = n3

	n4 := &ListNode{Val: 2}
	n3.Next = n4

	n5 := &ListNode{Val: 5}
	n4.Next = n5

	n6 := &ListNode{Val: 2}
	n5.Next = n6

	node := partition(n1, 3)
	log.Println(node)
}

func TestReverse(t *testing.T) {

	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n1.Next = n2
	n2.Next = n3

	n := reverseN(n1, 2)
	log.Println(n)
}

func TestReverseN(t *testing.T) {

	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5

	n := reverseN(n1, 3)
	log.Println("node:", n)

}

func TestReverseBetween(t *testing.T) {

	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n6 := &ListNode{Val: 6}
	n7 := &ListNode{Val: 7}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7

	node := reverseBetween(n1, 2, 6)
	log.Println("node:", node)

}

func TestReverseByNode(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n6 := &ListNode{Val: 6}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6

	node := reverseByNode(n1, n4)
	log.Println("node:", node)

}

func TestRevereKGroup(t *testing.T) {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n5 := &ListNode{Val: 5}
	n6 := &ListNode{Val: 6}
	n7 := &ListNode{Val: 7}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7

	node := reverseKGroup(n1, 3)
	log.Println(node)

}
