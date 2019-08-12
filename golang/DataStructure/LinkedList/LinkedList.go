package LinkedList

//ListNode defines a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//Length returns the length of the linked list instance
func (ll *ListNode) Length() int {
	len := 0
	for ll != nil {
		len++
		ll = ll.Next
	}
	return len
}

//Middle returns the middle point of the linked list instance
func (ll *ListNode) Middle() *ListNode {
	slow, fast := ll, ll
	for ll.Next != nil && ll.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
