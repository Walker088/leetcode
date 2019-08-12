package ConvertSortedListToBinaryTree

import (
	"fmt"
	"testing"

	ll "github.com/Walker088/leetcode/golang/DataStructure/LinkedList"
)

func Test_SortedListToBST(t *testing.T) {
	fmt.Println("Test_SortedListToBST")
	l5 := &ll.ListNode{Val: 9, Next: nil}
	l4 := &ll.ListNode{Val: 5, Next: l5}
	l3 := &ll.ListNode{Val: 0, Next: l4}
	l2 := &ll.ListNode{Val: -3, Next: l3}
	l1 := &ll.ListNode{Val: -10, Next: l2}
	fmt.Println("res: ", SortedListToBST(l1))
}
