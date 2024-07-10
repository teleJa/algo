package leetcode

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	dummyHead *Node
}

func (l MyLinkedList) AddAtHead(val int) {
	node := NewNode(val)
	h := l.dummyHead.next
	l.dummyHead.next = node
	node.next = h
}

func (l MyLinkedList) AddAtTail(val int) {
	node := NewNode(val)
	cur := l.dummyHead
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = node
}

func (l MyLinkedList) AddAtIndex(index, val int) {
	node := NewNode(val)
	cur := l.dummyHead
	i := 0
	for ; i < index; i++ {
		if cur.next == nil {
			break
		}
		cur = cur.next
	}
	if i == index {
		n := cur.next
		cur.next = node
		node.next = n
	}
}

func (l MyLinkedList) DeleteAtIndex(index int) {
	cur := l.dummyHead
	i := 0
	for ; i < index; i++ {
		if cur.next == nil {
			break
		}
		cur = cur.next
	}

	if i == index && cur.next != nil {
		n := cur.next
		cur.next = n.next
		n.next = nil
	}

}

func (l MyLinkedList) Get(index int) int {

	i := -1
	cur := l.dummyHead

	idx := 0
	for ; idx < index; idx++ {
		if cur.next == nil {
			break
		}
		cur = cur.next
	}

	if idx == index && cur.next != nil {
		return cur.next.val
	}
	return i
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		dummyHead: &Node{},
	}
}

func NewNode(val int) *Node {
	return &Node{
		val: val,
	}
}
