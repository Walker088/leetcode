package addTwoNumbers

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintFullList(l *ListNode) {
	for l.Next != nil {
		fmt.Println(l)
		l = l.Next
	}
	fmt.Println(l)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry, left := 0, 0
	var res *ListNode = &ListNode{Val: 0, Next: nil}
	var res_fix *ListNode = res
	for {
		if l1 != nil && l2 != nil {
			digit := l1.Val + l2.Val + carry
			left, carry = digit%10, digit/10
			res.Next = &ListNode{Val: left}
			l1, l2, res = l1.Next, l2.Next, res.Next
			continue
		}
		if l1 != nil || l2 != nil {
			if l1 != nil {
				digit := l1.Val + carry
				left, carry = digit%10, digit/10
				res.Next = &ListNode{Val: left}
				l1, res = l1.Next, res.Next
				continue
			}
			if l2 != nil {
				digit := l2.Val + carry
				left, carry = digit%10, digit/10
				res.Next = &ListNode{Val: left}
				l2, res = l2.Next, res.Next
				continue
			}
		}
		if l1 == nil && l2 == nil {
			if carry != 0 {
				res.Next = &ListNode{Val: 1}
			}
			break
		}
	}
	return res_fix.Next
}
