package BinaryTreePruning

import (
	"fmt"
	"testing"

	bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"
)

func Test_PruneTree(t *testing.T) {
	fmt.Println("Test_PruneTree")
	t4 := bt.NewBinaryTree(1, nil, nil)
	t3 := bt.NewBinaryTree(0, nil, nil)
	t2 := bt.NewBinaryTree(0, t3, t4)
	t1 := bt.NewBinaryTree(1, nil, t2)
	res := PruneTree(t1)
	res.WholeTreePrint("root", res.GetHeight())
}
