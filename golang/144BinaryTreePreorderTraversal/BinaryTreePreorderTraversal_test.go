package BinaryTreePreorder

import (
	"fmt"
	"testing"

	bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"
)

func Test_PreorderTraversal(t *testing.T) {
	fmt.Println("Test_PreorderTraversal")
	t4 := bt.NewBinaryTree(5, nil, nil)
	t5 := bt.NewBinaryTree(3, nil, nil)
	t2 := bt.NewBinaryTree(3, t4, t5)
	t7 := bt.NewBinaryTree(9, nil, nil)
	t3 := bt.NewBinaryTree(2, nil, t7)
	t1 := bt.NewBinaryTree(1, t2, t3)
	res := PreorderTraversal(t1)
	t1.PrintByPreOrder()
	fmt.Print(res)
}
