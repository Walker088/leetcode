package ConsBinaryTree

import (
	"fmt"

	bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"
)

func buildTree(postorder []int, ihm map[int]int, postStart int, postEnd int, inStart int) *bt.TreeNode {
	if postStart > postEnd {
		return nil
	}
	root := &bt.TreeNode{Val: postorder[postEnd]}
	rootID := ihm[root.Val]
	leftLen := rootID - inStart
	root.Left = buildTree(postorder, ihm, postStart, postStart+leftLen-1, inStart)
	root.Right = buildTree(postorder, ihm, postStart+leftLen, postEnd-1, rootID+1)
	return root
}

//BuildTree builds a binary tree by inorder/postorder traversal
func BuildTree(inorder []int, postorder []int) *bt.TreeNode {
	fmt.Println("BuildTree")
	if (len(inorder) != len(postorder)) || len(inorder) <= 0 {
		return nil
	}
	ihm := map[int]int{} //inorder hash map, records the index of each element of inorder list
	for i, v := range inorder {
		ihm[v] = i
	}
	return buildTree(postorder, ihm, 0, len(postorder)-1, 0)
}
