package tree

// 使用数组模拟完全二叉树实现一个小顶堆
type CBTree struct {
	nodes []int
}

// node为索引
func parent(node int) int {
	return node / 2
}

func left(node int) int {
	return node * 2
}

func right(node int) int {
	return node*2 + 1
}

func NeCBTree(size int) *CBTree {
	return &CBTree{
		// index=1为根节点
		nodes: make([]int, 1, size),
	}
}

func (t *CBTree) Push(node int) {
	// 先加到最后
	t.nodes = append(t.nodes, node)
	// 上浮
	t.swim(len(t.nodes) - 1)
}

func (t *CBTree) swim(index int) {
	cur := index
	pIndex := parent(cur)
	for t.nodes[pIndex] > t.nodes[cur] && pIndex > 0 {
		t.swap(cur, pIndex)
		cur = pIndex
		pIndex = parent(pIndex)
	}
}

func (t *CBTree) swap(i, j int) {
	t.nodes[i], t.nodes[j] = t.nodes[j], t.nodes[i]
}

func (t *CBTree) Pop() int {
	top := t.Peek()
	// 下沉
	t.sink(len(t.nodes) - 1)
	return top
}

func (t *CBTree) sink(index int) {
	t.nodes[1] = t.nodes[index]
	cur := 1
	min := cur
	l := left(cur)
	r := right(cur)
	for l < len(t.nodes) || r < len(t.nodes) {
		if l < len(t.nodes) && t.nodes[cur] > t.nodes[l] {
			cur = l
		}
		if r < len(t.nodes) && t.nodes[cur] > t.nodes[r] {
			cur = r
		}

		if min == cur {
			break
		}
		t.swap(min, cur)
		min = cur
		l = left(cur)
		r = right(cur)
	}
}

func (t *CBTree) sinkErr(index int) {
	t.nodes[1] = t.nodes[index]
	cur := 1
	min := cur
	l := left(cur)
	r := right(cur)
	for l < len(t.nodes) || r < len(t.nodes) {
		// 过早交互值,可能l对应的值比r对应的值还要小
		if l < len(t.nodes) && t.nodes[cur] > t.nodes[l] {
			t.swap(cur, l)
			cur = l
		}
		if r < len(t.nodes) && t.nodes[cur] > t.nodes[r] {
			t.swap(cur, r)
			cur = r
		}

		if min == cur {
			break
		}
		min = cur
		l = left(cur)
		r = right(cur)
	}
}

func (t *CBTree) Peek() int {
	return t.nodes[1]
}
