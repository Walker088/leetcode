package ConvertSortedListToBinaryTree

import (
	bt "github.com/Walker088/leetcode/golang/DataStructure/BinaryTree"
	ll "github.com/Walker088/leetcode/golang/DataStructure/LinkedList"
)

/*
SortedListToBST do the following job:
Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.
*/
func SortedListToBST(head *ll.ListNode) *bt.TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		//return &bt.TreeNode{head.Val, nil, nil}
		return bt.NewBinaryTree(head.Val, nil, nil)
	}
	return convHelper(head, nil)
}
func convHelper(head *ll.ListNode, tail *ll.ListNode) *bt.TreeNode {
	if head == tail {
		return nil
	}
	slow, fast := head, head
	for fast != tail && fast.Next != tail {
		slow = slow.Next
		fast = fast.Next.Next
	}
	res := bt.NewBinaryTree(slow.Val, nil, nil)
	//res := &bt.TreeNode{slow.Val, nil, nil}
	res.Left = convHelper(head, slow)
	res.Right = convHelper(slow.Next, tail)
	return res
}
