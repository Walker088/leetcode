package MaxBinTree

import bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"

func getMaxFromLst(nums []int) int {
	mi := 0
	for i, v := range nums {
		if nums[mi] < v {
			mi = i
		}
	}
	return mi
}

func constructMaximumBinaryTree(nums []int) *bt.TreeNode {
	//fmt.Println("constructMaximumBinaryTree")
	if len(nums) == 0 {
		return nil
	}
	mi := getMaxFromLst(nums)
	root := bt.TreeNode{Val: nums[mi]}
	root.Left = constructMaximumBinaryTree(nums[:mi])
	root.Right = constructMaximumBinaryTree(nums[mi+1:])
	return &root
}
