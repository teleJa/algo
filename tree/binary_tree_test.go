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

	// 前序遍历 // 入栈前
	//fmt.Println(root.Val)
	traverse(root.Left) // 理解为入栈
	// 中序遍历   // 出栈
	//fmt.Println(root.Val)
	traverse(root.Right) // 再入栈
	// 后序遍历
	fmt.Println(root.Val) // 出栈
}

func TestBFS(t *testing.T) {

	//bfs1(root)
	//bfs2(root)
	bfs3(root)
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

type State struct {
	Node  *TreeNode
	Depth int
}

func bfs3(root *TreeNode) {

	var queue []State
	if root != nil {
		queue = append(queue, State{root, 1})
	}
	// depth 增量
	delta := 1
	for len(queue) > 0 {
		st := queue[0]
		log.Println("node:", st.Node.Val, "depth:", st.Depth)

		queue = queue[1:]
		if st.Node.Left != nil {
			left := State{
				Node:  st.Node.Left,
				Depth: st.Depth + delta,
			}
			queue = append(queue, left)
		}

		if st.Node.Right != nil {
			right := State{
				Node:  st.Node.Right,
				Depth: st.Depth + delta,
			}
			queue = append(queue, right)
		}

	}

}
