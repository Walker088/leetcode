package CheckBinaryTreeCompleteness

import bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"

//IsCompleteTree returns whether the input treenode is a complete binary tree
func IsCompleteTree(root *bt.TreeNode) bool {
	node := &bt.TreeNode{}
	queue := []*bt.TreeNode{}
	queue = append(queue, root)

	for {
		node = queue[0]
		queue = queue[1:]
		if node.Left == nil {
			if node.Right != nil {
				return false
			}
			break
		}
		queue = append(queue, node.Left)
		if node.Right == nil {
			break
		}
		queue = append(queue, node.Right)
	}
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		if node.Left != nil || node.Right != nil {
			return false
		}
	}
	return true
}
