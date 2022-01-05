package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3, t3 *ListNode
	t1 := l1
	t2 := l2

	carry := 0
	unit := 0
	for t1 != nil || t2 != nil {
		var v1, v2 int

		if t1 == nil {
			v1 = 0
		} else {
			v1 = t1.Val
		}

		if t2 == nil {
			v2 = 0
		} else {
			v2 = t2.Val
		}

		sum := v1 + v2 + carry

		if sum > 9 {
			unit = sum % 10
			carry = sum / 10
		} else {
			unit = sum
			carry = 0

		}

		if l3 == nil {
			t3 = &ListNode{Val: unit, Next: nil}
			l3 = t3
		} else {
			t3.Next = &ListNode{Val: unit, Next: nil}
			t3 = t3.Next
		}

		if t1 != nil {
			t1 = t1.Next
		}

		if t2 != nil {
			t2 = t2.Next
		}
	}

	if carry > 0 {
		t3.Next = &ListNode{Val: carry, Next: nil}
	}
	return l3
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	l3 := addTwoNumbers(l1, l2)

	t1 := l1
	t2 := l2
	t3 := l3

	for t1 != nil {
		fmt.Print(t1.Val, " ")
		t1 = t1.Next
	}
	fmt.Println()

	for t2 != nil {
		fmt.Print(t2.Val, " ")
		t2 = t2.Next
	}
	fmt.Println()

	for t3 != nil {
		fmt.Print(t3.Val, " ")
		t3 = t3.Next
	}
	fmt.Println()
}
