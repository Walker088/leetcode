package ConvertSortedListToBinaryTree

import (
	"fmt"
	"testing"

	ll "github.com/Walker088/leetcode/golang/DataStructure/LinkedList"
)

func Test_SortedListToBST(t *testing.T) {
	fmt.Println("Test_SortedListToBST")
	l5 := &ll.ListNode{9, nil}
	l4 := &ll.ListNode{5, l5}
	l3 := &ll.ListNode{0, l4}
	l2 := &ll.ListNode{-3, l3}
	l1 := &ll.ListNode{-10, l2}
	fmt.Println("res: ", SortedListToBST(l1))
}
