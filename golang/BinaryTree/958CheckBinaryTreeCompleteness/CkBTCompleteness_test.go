package CheckBinaryTreeCompleteness

import (
	"fmt"
	"testing"

	bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"
)

func Test_IsCompleteTree(t *testing.T) {
	fmt.Println("Test_IsCompleteTree")
	t4 := bt.NewBinaryTree(5, nil, nil)
	t5 := bt.NewBinaryTree(3, nil, nil)
	t2 := bt.NewBinaryTree(3, t4, t5)
	t7 := bt.NewBinaryTree(9, nil, nil)
	t3 := bt.NewBinaryTree(2, nil, t7)
	t1 := bt.NewBinaryTree(1, t2, t3)
	res := IsCompleteTree(t1)
	//t1.WholeTreePrint("Root", t1.GetHeight())
	fmt.Println("t1 is a complete binary tree: ", res)
}
