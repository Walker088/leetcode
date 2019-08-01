package RemoveNthFromEnd

//ListNode definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

//Length defines
func Length(l *ListNode) int {
	if l == nil {
		return 0
	}
	len := 1
	for l.Next != nil {
		len++
		l = l.Next
	}
	return len
}

//RemoveElement removes the n-th node from l
func RemoveElement(l *ListNode, n int) *ListNode {
	if l.Next == nil {
		return nil
	}
	for i := 0; i <= n; i++ {
		if i == n-1 {
			l.Next = l.Next.Next
			break
		}
		l = l.Next
	}
	return l
}

//RemoveNthFromEnd remove the n-th node from tail
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	//fmt.Println("RemoveNthFromEnd")
	if head == nil {
		return nil
	}
	length := Length(head)
	target := length - n + 1
	dummy := &ListNode{0, head}
	if n > length {
		return nil
	}
	RemoveElement(dummy, target)
	return dummy.Next
}
