package DistributeCoins

import (
	"fmt"
	"testing"

	bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"
)

func Test_DistributeCoins(t *testing.T) {
	fmt.Println("Test_DistributeCoins")
	t2 := bt.NewBinaryTree(0, nil, nil)
	t3 := bt.NewBinaryTree(0, nil, nil)
	t1 := bt.NewBinaryTree(3, t2, t3)
	res := DistributeCoins(t1)
	fmt.Println("res: ", res)

	t4 := bt.NewBinaryTree(3, nil, nil)
	t2 = bt.NewBinaryTree(0, nil, t4)
	t3 = bt.NewBinaryTree(0, nil, nil)
	t1 = bt.NewBinaryTree(1, t2, t3)
	res = DistributeCoins(t1)
	fmt.Println("res: ", res)
}
