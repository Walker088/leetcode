package BinaryTreeInorder

//TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//InorderTraversal returns
func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	it := InorderTraversal(root.Left)
	it = append(it, root.Val)
	it = append(it, InorderTraversal(root.Right)...)
	return it
}
