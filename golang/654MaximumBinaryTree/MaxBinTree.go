package MaxBinTree

//TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMaxFromLst(nums []int) int {
	mi := 0
	for i, v := range nums {
		if nums[mi] < v {
			mi = i
		}
	}
	return mi
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	//fmt.Println("constructMaximumBinaryTree")
	if len(nums) == 0 {
		return nil
	}
	mi := getMaxFromLst(nums)
	root := TreeNode{Val: nums[mi]}
	root.Left = constructMaximumBinaryTree(nums[:mi])
	root.Right = constructMaximumBinaryTree(nums[mi+1:])
	return &root
}
