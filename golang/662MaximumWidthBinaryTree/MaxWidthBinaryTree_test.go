package MaxWidthBinaryTree

import (
	"fmt"
	"testing"
)

func getHeight(n *TreeNode) int {
	if n == nil {
		return 0
	}
	lefth := getHeight(n.Left)
	righth := getHeight(n.Right)
	if lefth > righth {
		return lefth + 1
	}
	return righth + 1
}

// Left-Root-Right
func inorder(n *TreeNode) {
	if n == nil {
		return
	}
	inorder(n.Left)
	fmt.Print(n.Val, " ")
	inorder(n.Right)
}

// Root-Left-Right
func preorder(n *TreeNode) {
	if n == nil {
		return
	}
	fmt.Print(n.Val, " ")
	preorder(n.Left)
	preorder(n.Right)
}

// Left-Right-Root
func postorder(n *TreeNode) {
	if n == nil {
		return
	}
	postorder(n.Left)
	postorder(n.Right)
	fmt.Print(n.Val, " ")
}

func print(n *TreeNode, s string, h int) {
	if n == nil {
		return
	}
	fmt.Println(s, ": ", n.Val, " height: ", h)
	print(n.Left, "L", h-1)
	print(n.Right, "R", h-1)
}

func Test_widthOfBinaryTree(t *testing.T) {
	fmt.Println("Test_widthOfBinaryTree")
	t4 := &TreeNode{5, nil, nil}
	t5 := &TreeNode{3, nil, nil}
	t2 := &TreeNode{3, t4, t5}
	t7 := &TreeNode{9, nil, nil}
	t3 := &TreeNode{2, nil, t7}
	t1 := &TreeNode{1, t2, t3}
	//h := getHeight(t1)
	//print(t1, "Root", h)
	//preorder(t1)
	res := widthOfBinaryTree(t1)
	fmt.Println("res: ", res)
}
