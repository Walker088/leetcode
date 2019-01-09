package addTwoNumbers

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	var res = &ListNode{}
	var l3 = &ListNode{Val: 1, Next: nil}
	var l2 = &ListNode{Val: 5, Next: l3}
	var l1 = &ListNode{Val: 4, Next: l2}

	var lc = &ListNode{Val: 2, Next: nil}
	var lb = &ListNode{Val: 9, Next: lc}
	var la = &ListNode{Val: 3, Next: lb}

	res = addTwoNumbers(la, l1)
	if res != nil {
		t.Log("test passed!!!")
	}
	t.Log("test passed!!!")
}

func Test_addTwoNumbers2(t *testing.T) {
	var res = &ListNode{}

	var l1 = &ListNode{Val: 5, Next: nil}
	var l2 = &ListNode{Val: 5, Next: nil}

	res = addTwoNumbers(l1, l2)
	if sum := (res.Val + 10*(res.Next.Val)); sum != 10 {
		t.Error("Test Faild!, res suppose to be 10, but got:", sum)
	}
	t.Log("test passed!!!")
}

func Test_addTwoNumbers3(t *testing.T) {
	var res = &ListNode{}

	var l3 = &ListNode{Val: 9, Next: nil}
	var l2 = &ListNode{Val: 5, Next: l3}
	var l1 = &ListNode{Val: 5, Next: l2}

	var la = &ListNode{Val: 5, Next: nil}

	res = addTwoNumbers(l1, la)
	if sum := (res.Val + 10*(res.Next.Val) + 100*(res.Next.Next.Val)); sum != 960 {
		t.Error("Test Faild!, res suppose to be 960, but got:", sum)
	}
	t.Log("test passed!!!")
}
