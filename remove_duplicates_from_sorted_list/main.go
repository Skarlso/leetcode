package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// The list is guaranteed to be sorted in ascending order.
// So just skip if the next one is equal to the prev one.
func deleteDuplicates(head *ListNode) *ListNode {
	prev := head
	current := head

	for current != nil {
		if current.Val == prev.Val {
			prev.Next = current.Next
		} else {
			prev = current
		}
		current = current.Next
	}

	return head
}
