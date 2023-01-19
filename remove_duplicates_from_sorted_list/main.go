package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// deal with deleting every value at the begin of the list
	// that would start with the value that is to be deleted.
	// [7, 7, 7, 7]
	for head != nil && head.Val == val {
		head = head.Next
	}

	prev := &ListNode{}
	current := head

	// Normal delete operation after that.
	for current != nil {
		if current.Val == val {
			prev.Next = current.Next
		} else {
			prev = current
		}
		current = current.Next
	}

	return head
}
