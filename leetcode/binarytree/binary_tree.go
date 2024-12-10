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

// 分解
func invertTree2(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	// 翻转左子树
	root.Left = invertTree2(root.Left)
	// 翻转右子树
	root.Right = invertTree2(root.Right)

	root.Left, root.Right = root.Right, root.Left
	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 116 填充每一个节点的下一个右侧节点指针
// 每次遍历2个节点即可解决兄弟节点之间的关系,对于n叉树,可拓展为每次遍历时传递该层的全部节点
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var f func(left, right *Node)
	f = func(node1, node2 *Node) {
		if node1 == nil {
			return
		}
		node1.Next = node2
		f(node1.Left, node1.Right)
		f(node1.Right, node2.Left)
		f(node2.Left, node2.Right)
	}
	f(root.Left, root.Right)
	return root
}
