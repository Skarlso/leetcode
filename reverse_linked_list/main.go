package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// The trivial thing to do would be to construct a new list.
// Loop through the list and construct it backwards. But how,
// since we loop in forwarding order.
// I could construct a list from the values and then build up a reversed list.
// That sounds terribly inefficient.
// Update the reference to point to Prev but before that, get Next.
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var (
		prev, next *ListNode
	)
	current := head
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
