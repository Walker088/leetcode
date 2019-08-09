package BinaryTreePreorder

import bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"

//PreorderTraversal returns preorder traversal of a binary tree
func PreorderTraversal(root *bt.TreeNode) []int {
	//fmt.Println("PreorderTraversal")
	if root == nil {
		return nil
	}
	Lst := []int{root.Val}
	Lst = append(Lst, PreorderTraversal(root.Left)...)
	Lst = append(Lst, PreorderTraversal(root.Right)...)
	return Lst
}
