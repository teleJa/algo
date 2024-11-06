package tree

import (
	"fmt"
	"testing"
)

func TestTraverse(t *testing.T) {

	node7 := &TreeNode{Val: 7}
	node6 := &TreeNode{Val: 6}
	node5 := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 4}
	node3 := &TreeNode{Val: 3, Left: node5, Right: node6}
	node2 := &TreeNode{Val: 2, Left: node7, Right: node4}
	root := &TreeNode{Val: 1, Left: node2, Right: node3}

	traverse(root)
}

func traverse(root *TreeNode) {

	if root == nil {
		return
	}

	// 前序遍历
	//fmt.Println(root.Val)
	traverse(root.Left)
	// 中序遍历
	fmt.Println(root.Val)
	traverse(root.Right)
	// 后序遍历
	//fmt.Println(root.Val)
}
