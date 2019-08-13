package BinaryTreeRightSideView

import bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"

//RightSideView returns a list of the right side view of a binary tree
func RightSideView(root *bt.TreeNode) []int {
	res := []int{}
	mkRVLst(root, &res, 0)
	return res
}
func mkRVLst(root *bt.TreeNode, lst *[]int, depth int) {
	if root == nil {
		return
	}
	if depth == len(*lst) {
		*lst = append(*lst, root.Val)
	}
	mkRVLst(root.Right, lst, depth+1)
	mkRVLst(root.Left, lst, depth+1)
}
