package binarytree

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 104 二叉树的最大深度
func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	depth := 0
	var queue []*TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		depth++
		sz := len(queue)
		for i := 0; i < sz; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[sz:]
	}
	return depth
}
