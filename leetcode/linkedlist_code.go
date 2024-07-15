package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 21
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	dummy := &ListNode{}
	p := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}

	if list1 != nil {
		p.Next = list1
	}

	if list2 != nil {
		p.Next = list2
	}

	return dummy.Next
}

// 86
func partition(head *ListNode, x int) *ListNode {

	dummy1 := &ListNode{}
	dummy2 := &ListNode{}

	p1 := dummy1
	p2 := dummy2

	for head != nil {
		if head.Val < x {
			p1.Next = head
			p1 = p1.Next
		} else {
			p2.Next = head
			p2 = p2.Next
		}

		// 断开原链表,避免出现环
		n := head.Next
		head.Next = nil
		head = n
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}
