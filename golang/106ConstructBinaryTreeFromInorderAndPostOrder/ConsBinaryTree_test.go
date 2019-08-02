package ConsBinaryTree

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
func print(n *TreeNode, s string, h int) {
	if n == nil {
		return
	}
	fmt.Println(s, ": ", n.Val, " height: ", h)
	print(n.Left, "L", h-1)
	print(n.Right, "R", h-1)
}

func Test_BuildTree(t *testing.T) {
	fmt.Println("Test_BuildTree")
	io := []int{9, 3, 15, 20, 7}
	po := []int{9, 15, 7, 20, 3}
	//var test []int
	//fmt.Println("len(test)", len(test))
	bt := BuildTree(io, po)
	fmt.Println("Binary Tree: ", bt)
	fmt.Println("Whole Tree:")
	print(bt, "Root", getHeight(bt))
}
