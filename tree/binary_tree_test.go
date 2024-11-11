package tree

import (
	"fmt"
	"log"
	"testing"
)

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

func TestDFS(t *testing.T) {
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
	//fmt.Println(root.Val)
	traverse(root.Right)
	// 后序遍历
	fmt.Println(root.Val)
}

func TestBFS(t *testing.T) {

	//bfs1(root)
	bfs2(root)
}

// 层序遍历1
func bfs1(root *TreeNode) {

	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Println(node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

}

// 打印深度
func bfs2(root *TreeNode) {

	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}

	depth := 1
	for len(queue) > 0 {

		sz := len(queue)
		for i := 0; i < sz; i++ {
			node := queue[i]

			log.Println("node:", node.Val, "depth:", depth)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[sz:]
		depth++
	}

}
