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

// 19
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1 := head
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}

	p2 := head
	dummy := &ListNode{}
	p := dummy
	for p1 != nil {
		p.Next = p2
		p = p.Next
		p1 = p1.Next
		p2 = p2.Next
	}
	p.Next = p2.Next
	return dummy.Next
}

// 19
func removeNthFromEnd2(head *ListNode, n int) *ListNode {

	dummy := &ListNode{}
	dummy.Next = head
	// 寻找倒数第n+1个节点
	node := FindLastNode(dummy, n+1)
	node.Next = node.Next.Next
	return dummy.Next
}

func FindLastNode(head *ListNode, n int) *ListNode {
	p1 := head
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}

	p2 := head
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}

// 876
func middleNode(head *ListNode) *ListNode {

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 141
func hasCycle(head *ListNode) bool {

	slow := head
	fast := head
	has := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			has = true
			break
		}
	}
	return has
}

// 142
func detectCycle(head *ListNode) *ListNode {

	slow := head
	fast := head
	has := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			has = true
			break
		}
	}

	if !has {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
