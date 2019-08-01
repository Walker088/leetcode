package RemoveNthFromEnd

import (
	"fmt"
	"testing"
)

func makeNodeLst(nodes ...*ListNode) *ListNode {
	for i, n := range nodes {
		if i != len(nodes)-1 {
			n.Val = i
			n.Next = nodes[(i+1)%len(nodes)]
		} else {
			n.Val = i
			n.Next = nil
		}
		//fmt.Println(i, n)
	}
	return nodes[0]
}

func DisPlay(l *ListNode) {
	for l != nil {
		fmt.Printf("%v ->", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func Test_RemoveNthFromEnd(t *testing.T) {
	var (
		n5 = &ListNode{5, nil}
		n4 = &ListNode{4, n5}
		n3 = &ListNode{3, n4}
		n2 = &ListNode{2, n3}
		n1 = &ListNode{1, n2}
	)
	DisPlay(n1)
	res1 := RemoveNthFromEnd(n1, 3)
	DisPlay(res1)

	fmt.Println("========================")

	DisPlay(n5)
	res2 := RemoveNthFromEnd(n5, 1)
	DisPlay(res2)
}
