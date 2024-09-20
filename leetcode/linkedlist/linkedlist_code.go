package linkedlist

import "math"

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

// 19 双指针法 假设链表长度为k,到达倒数第n个节点需要走k-n步
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// p1先走n步
	p1 := head
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}

	// p2从头开始走 k-n步
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

// 876 寻找中间节点
func middleNode(head *ListNode) *ListNode {

	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 141快慢指针相遇即为有环
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

	// 假设相遇点为m(必然在环中),慢指针走了k步,快指针必然走了2k步,则k是环的长度
	// 此时从慢指针重新走头节点出发走k-m步,快指针每次走一步也从相遇点出发,再次相遇必然就是环的起点

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

// 160 获取相交节点,假设x为相交的链表部分
// 则整合为一个链表后为 A + x + B = B + x + A
// 即对于2个指针,分别遍历完整合后的链表只要相等必然相交
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
	return prev
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
		//此时的n即为n-m+1,即为反转前n个节点
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

// 反转完毕后a是尾节点
// 返回的链表是原链表的一部分
func reverseByNode(a *ListNode, b *ListNode) *ListNode {
	var prev *ListNode
	for a != b {
		n := a.Next
		a.Next = prev
		prev = a
		a = n
	}
	return prev
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}

	n := reverseByNode(a, b)
	a.Next = reverseKGroup(b, k)
	return n
}

// 83 删除排序链表中的重复元素
// 有序的链表,重复的元素必然是相连的
func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast != nil {
		if slow.Val != fast.Val {
			slow = slow.Next
			slow.Val = fast.Val
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head
	for fast != nil {
		if slow.Val != fast.Val {
			// 直接把slow.next指向当前节点
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}

// 234 回文链表
func isPalindrome(head *ListNode) bool {

	// 复制链表
	dummy := &ListNode{}
	p := dummy
	h0 := head
	for h0 != nil {
		p.Next = &ListNode{Val: h0.Val}
		h0 = h0.Next
		p = p.Next
	}

	newHead := reverseList(head)
	h := dummy.Next
	for h != nil {
		if h.Val != newHead.Val {
			return false
		}
		h = h.Next
		newHead = newHead.Next
	}
	return true
}

// 88. 合并两个有序数组
func merge(nums1 []int, m int, nums2 []int, n int) {

	// 从后向前合并(nums1数组后面是空的)
	m0, n0 := m-1, n-1
	p := m + n - 1
	for m0 >= 0 && n0 >= 0 {
		if nums1[m0] < nums2[n0] {
			nums1[p] = nums2[n0]
			n0--
		} else {
			nums1[p] = nums1[m0]
			m0--
		}
		p--
	}

	// 判断num2是否遍历完成
	for n0 >= 0 {
		nums1[p] = nums2[n0]
		n0--
		p--
	}
}

// 977. 有序数组的平方
// nums中可能存在负数,平方后变成正数
func sortedSquares(nums []int) []int {

	// m0 代表左侧小于0数据 n0代表右侧大于0数据
	m0, n0 := 0, 0
	for i := range nums {
		if nums[i] >= 0 {
			n0 = i
			break
		} else {
			n0++
		}
	}

	m0 = n0 - 1
	res := make([]int, len(nums))
	index := 0
	// 合并数组
	for m0 > -1 && n0 < len(nums) {
		abs := math.Abs(float64(nums[m0]))
		if abs < float64(nums[n0]) {
			res[index] = int(abs)
			m0--
		} else {
			res[index] = nums[n0]
			n0++
		}
		index++
	}

	for m0 > -1 {
		abs := math.Abs(float64(nums[m0]))
		res[index] = int(abs)
		m0--
		index++
	}

	for n0 < len(nums) {
		res[index] = nums[n0]
		n0++
		index++
	}

	for i := range res {
		res[i] = res[i] * res[i]
	}
	return res

}

// 双指针初始化在数组的尾部,从绝对值来看,左侧负数子数组从左向右是递减,右侧正数子数组从右向左也是递减(v字形)
func sortedSquares2(nums []int) []int {

	// 此时left,right 代表2个子数组绝对值最大的索引
	left, right := 0, len(nums)-1
	res := make([]int, len(nums))
	p := len(nums) - 1
	for left <= right {
		l := math.Abs(float64(nums[left]))
		r := math.Abs(float64(nums[right]))
		if l > r {
			res[p] = int(l) * int(l)
			left++
		} else {
			res[p] = int(r) * int(r)
			right--
		}
		p--
	}
	return res
}

// 360 有序转化数组
// 给你一个已经 排好序 的整数数组 nums 和整数 a 、 b 、 c 。对于数组中的每一个元素 nums[i] ，计算函数值 f(x) = ax2 + bx + c ，请 按升序返回数组 。
// 此函数为抛物线,a > 0 开口向上, a < 0 开口向下
// a=0的时候考虑 b的情况 b > 0,f(x)逐渐递增(退化为开口向下抛物线的左侧), b <0 ,f(x)逐渐递减(退化为开口向下抛物线的右侧),因此 a=0应归为抛物线开口向下的情况
func sortTransformedArray(nums []int, a, b, c int) []int {

	left, right := 0, len(nums)-1
	p := 0
	if a > 0 {
		p = right
	}

	fn := func(x, a, b, c int) int {
		return a*x*x + b*x + c
	}

	res := make([]int, len(nums))
	for left <= right {

		l0 := fn(nums[left], a, b, c)
		r0 := fn(nums[right], a, b, c)
		// 开口向上,两边向中心递减
		if a > 0 {
			if l0 > r0 {
				res[p] = l0
				p--
				left++
			} else {
				res[p] = r0
				p--
				right--
			}
		} else {
			if l0 > r0 {
				res[p] = r0
				right--
				p++
			} else {
				res[p] = l0
				left++
				p++
			}
		}

	}
	return res
}
