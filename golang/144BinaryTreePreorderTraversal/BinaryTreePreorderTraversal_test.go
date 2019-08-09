package BinaryTreePreorder

import (
	"fmt"
	"testing"
)

// Root-Left-Right
func preorder(n *TreeNode) {
	if n == nil {
		return
	}
	fmt.Print(n.Val, " ")
	preorder(n.Left)
	preorder(n.Right)
}

func Test_PreorderTraversal(t *testing.T) {
	fmt.Println("Test_PreorderTraversal")
	t4 := &TreeNode{5, nil, nil}
	t5 := &TreeNode{3, nil, nil}
	t2 := &TreeNode{3, t4, t5}
	t7 := &TreeNode{9, nil, nil}
	t3 := &TreeNode{2, nil, t7}
	t1 := &TreeNode{1, t2, t3}
	res := PreorderTraversal(t1)
	preorder(t1)
	fmt.Print(res)
}
