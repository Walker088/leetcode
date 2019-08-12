package CompleteBinaryTreeInserter

import (
	"fmt"
	"testing"

	bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"
)

func Test_Constructor(t *testing.T) {
	fmt.Println("Test_Constructor")
	t4 := bt.NewBinaryTree(5, nil, nil)
	t5 := bt.NewBinaryTree(3, nil, nil)
	t2 := bt.NewBinaryTree(3, t4, t5)
	t7 := bt.NewBinaryTree(9, nil, nil)
	t3 := bt.NewBinaryTree(2, nil, t7)
	t1 := bt.NewBinaryTree(1, t2, t3)
	cbti := Constructor(t1) //CBTInserter
	fmt.Println("CBTInserter: ", cbti)
	for _, v := range cbti.treeLst {
		fmt.Println(v.Val)
	}
}
