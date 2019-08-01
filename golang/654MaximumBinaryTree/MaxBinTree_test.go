package MaxBinTree

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

func Test_constructMaximumBinaryTree(t *testing.T) {
	fmt.Println("Test_constructMaximumBinaryTree")
	testCase := []int{3, 2, 1, 6, 0, 5}
	res := constructMaximumBinaryTree(testCase)
	h := getHeight(res)
	fmt.Println("res: ", *res, "height of tree: ", h)
	print(res, "Root", h)
}
