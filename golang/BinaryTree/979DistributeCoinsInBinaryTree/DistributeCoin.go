package DistributeCoins

import (
	bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//DistributeCoins returns the number of moves required to make every node have exactly one coin.
func DistributeCoins(root *bt.TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Left != nil {
		res += DistributeCoins(root.Left)
		res += abs(root.Left.Val - 1)
		root.Val += root.Left.Val - 1
	}
	if root.Right != nil {
		res += DistributeCoins(root.Right)
		res += abs(root.Right.Val - 1)
		root.Val += root.Right.Val - 1
	}
	return res
}
