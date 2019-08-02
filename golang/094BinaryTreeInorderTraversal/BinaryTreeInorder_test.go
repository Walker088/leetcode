package BinaryTreeInorder

import (
	"fmt"
	"testing"
)

// Left-Root-Right
func inorder(n *TreeNode) {
	if n == nil {
		return
	}
	inorder(n.Left)
	fmt.Print(n.Val, " ")
	inorder(n.Right)
}

func print(n *TreeNode, s string, h int) {
	if n == nil {
		return
	}
	fmt.Println(s, ": ", n.Val, " height: ", h)
	print(n.Left, "L", h-1)
	print(n.Right, "R", h-1)
}

func Test_InorderTraversal(t *testing.T) {
	fmt.Println("Test_InorderTraversal")
	t4 := &TreeNode{5, nil, nil}
	t5 := &TreeNode{3, nil, nil}
	t2 := &TreeNode{3, t4, t5}
	t7 := &TreeNode{9, nil, nil}
	t3 := &TreeNode{2, nil, t7}
	t1 := &TreeNode{1, t2, t3}
	res := InorderTraversal(t1)
	inorder(t1)
	fmt.Print(res)
}
