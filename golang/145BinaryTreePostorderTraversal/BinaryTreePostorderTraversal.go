package BinaryTreePostorder

//TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//PostorderTraversal returns preorder traversal of a binary tree
func PostorderTraversal(root *TreeNode) []int {
	//fmt.Println("PostorderTraversal")
	if root == nil {
		return nil
	}
	Lst := PostorderTraversal(root.Left)
	Lst = append(Lst, PostorderTraversal(root.Right)...)
	Lst = append(Lst, root.Val)
	return Lst
}
