package BinaryTreePruning

import bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"

//PruneTree returns
func PruneTree(root *bt.TreeNode) *bt.TreeNode {
	if root == nil {
		return nil
	}
	root.Left = PruneTree(root.Left)
	root.Right = PruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}
