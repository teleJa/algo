package binarytree

import "testing"

var (
	node7 = &TreeNode{Val: 7}
	node6 = &TreeNode{Val: 6}
	node5 = &TreeNode{Val: 5}
	node4 = &TreeNode{Val: 4}
	node3 = &TreeNode{Val: 3, Left: node6, Right: node7}
	node2 = &TreeNode{Val: 2, Left: node4, Right: node5}
	root  = &TreeNode{Val: 1, Left: node2, Right: node3}

	//     1
	//  2	  3
	// 4  5   6 7
)

func Test_flatten(t *testing.T) {
	flatten(root)
}
