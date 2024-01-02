package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 */
// Naive solution.
func hasCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	m := make(map[*ListNode]struct{})
	next := head
	for next != nil {
		if _, ok := m[next]; ok {
			return next
		}
		m[next] = struct{}{}
		next = next.Next
	}
	return nil
}

/*
Intuition- let's say 2 people are running in a circular track, one person is running slowly and another person is running faster(2 times the speed of first person)

After a certain period of time person 2 again meet or overtake person 1,

In that case we can conclude that the track is circular ( replace running track with our Linked list)

// I got to know it from the comment section of You tube, so credit goes to commenter
*/
