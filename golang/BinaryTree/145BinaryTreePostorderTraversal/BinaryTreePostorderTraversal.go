package BinaryTreePostorder

import bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"

//PostorderTraversal returns preorder traversal of a binary tree
func PostorderTraversal(root *bt.TreeNode) []int {
	//fmt.Println("PostorderTraversal")
	if root == nil {
		return nil
	}
	Lst := PostorderTraversal(root.Left)
	Lst = append(Lst, PostorderTraversal(root.Right)...)
	Lst = append(Lst, root.Val)
	return Lst
}
