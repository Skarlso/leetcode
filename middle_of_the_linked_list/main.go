package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 */
// Naive solution.
func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// Again, think about fast and slow pointer... If the fast pointer goes twice as fast and reaches the end
// the SLOW will be the middle one!!
