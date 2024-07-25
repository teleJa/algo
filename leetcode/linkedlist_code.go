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

// 86 拆分为2个小链表后合并
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

// 160
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else if p2 == nil {
			p2 = headA
		} else {
			p1 = p1.Next
			p2 = p2.Next
		}
	}
	return p1
}

// 206
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return last
}

// 206 迭代
func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		n := cur.Next
		cur.Next = prev
		prev = cur
		cur = n
	}
	return cur
}

// 使用+1的方式计数需要多传递一个变量
/*func reverseN(head *ListNode, n int) *ListNode {
	return reverseN0(head, 1, n)
}

var nextNode *ListNode

func reverseN0(head *ListNode, i, n int) *ListNode {
	if i == n {
		nextNode = head.Next
		return head
	}
	last := reverseN0(head.Next, i+1, n)
	head.Next.Next = head
	head.Next = nextNode
	return last
}*/

var nextNode *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		nextNode = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = nextNode
	return last
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == 1 {
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n)
	return head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	h := reverseN(head, k)
	cur, prev := nextNode, nextNode
	for i := 1; i < k; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}

	prev.Next = reverseKGroup(nextNode, k)
	return h
}
