package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// The list is guaranteed to be sorted in ascending order.
// So just skip if the next one is equal to the prev one.
func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil {
		for current.Next != nil && current.Val == current.Next.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}

	return head
}
