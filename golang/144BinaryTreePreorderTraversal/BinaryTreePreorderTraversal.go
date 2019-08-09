package BinaryTreePreorder

//TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//PreorderTraversal returns preorder traversal of a binary tree
func PreorderTraversal(root *TreeNode) []int {
	//fmt.Println("PreorderTraversal")
	if root == nil {
		return nil
	}
	Lst := []int{root.Val}
	Lst = append(Lst, PreorderTraversal(root.Left)...)
	Lst = append(Lst, PreorderTraversal(root.Right)...)
	return Lst
}
