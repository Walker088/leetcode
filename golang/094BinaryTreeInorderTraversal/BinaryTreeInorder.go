package BinaryTreeInorder

import bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"

//InorderTraversal returns
func InorderTraversal(root *bt.TreeNode) []int {
	if root == nil {
		return nil
	}
	it := InorderTraversal(root.Left)
	it = append(it, root.Val)
	it = append(it, InorderTraversal(root.Right)...)
	return it
}
