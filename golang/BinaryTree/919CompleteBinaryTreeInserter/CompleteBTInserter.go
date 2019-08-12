package CompleteBinaryTreeInserter

import bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"

//CBTInserter defines the struct of Complete Binary Tree
type CBTInserter struct {
	treeLst []*bt.TreeNode
}

//Constructor returns an instance of CBTInserter
func Constructor(root *bt.TreeNode) CBTInserter {
	return CBTInserter{treeLst: mkbfsLst(root)}
}
func mkbfsLst(root *bt.TreeNode) []*bt.TreeNode {
	var queue []*bt.TreeNode
	var resp []*bt.TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		resp = append(resp, node)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return resp
}

//Insert inserts an integer to the tree
func (ci *CBTInserter) Insert(v int) int {
	N := len(ci.treeLst)
	node := &bt.TreeNode{Val: v}
	ci.treeLst = append(ci.treeLst, node)
	if N%2 == 1 {
		ci.treeLst[(N-1)/2].Left = node
	} else {
		ci.treeLst[(N-1)/2].Right = node
	}
	return ci.treeLst[(N-1)/2].Val
}

//Getroot returns the root
func (ci *CBTInserter) Getroot() *bt.TreeNode {
	return ci.treeLst[0]
}
