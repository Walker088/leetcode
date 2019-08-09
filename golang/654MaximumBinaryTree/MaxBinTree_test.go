package MaxBinTree

import (
	"fmt"
	"testing"
)

func Test_constructMaximumBinaryTree(t *testing.T) {
	fmt.Println("Test_constructMaximumBinaryTree")
	testCase := []int{3, 2, 1, 6, 0, 5}
	res := constructMaximumBinaryTree(testCase)
	h := res.GetHeight()
	fmt.Println("res: ", *res, "height of tree: ", h)
	print(res, "Root", h)
}
