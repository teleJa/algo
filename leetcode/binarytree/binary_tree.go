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

// 144 二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	var res []int
	preorderTraversal0(root, &res)
	return res
}

func preorderTraversal0(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	preorderTraversal0(root.Left, res)
	preorderTraversal0(root.Right, res)
}

// 226 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {

	var f func(n *TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		root.Left, root.Right = root.Right, root.Left
		f(root.Left)
		f(root.Right)
	}
	f(root)
	return root
}
