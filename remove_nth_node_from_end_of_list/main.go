package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// dummy node solution
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// seek to the end of the list
	dummy := &ListNode{Next: head}
	fast := head
	slow := dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// distance between fast and slow should be n
	// make slow point to next next and return head.
	//
	slow.Next = slow.Next.Next

	return dummy.Next
}

// My solution:
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	// seek to the end of the list
//	numberOfNodes := 1
//	last := head
//	for {
//		if last.Next == nil {
//			break
//		}
//
//		numberOfNodes++
//		last = last.Next
//	}
//
//	if numberOfNodes == 1 && n == 1 {
//		return nil
//	}
//
//	if numberOfNodes <= 1 {
//		return head
//	}
//
//	if numberOfNodes == 2 && n == 2 {
//		return head.Next
//	}
//
//	current := head
//	for i := 1; i < numberOfNodes-n; i++ {
//		current = current.Next
//	}
//
//	if current.Next != nil {
//		current.Next = current.Next.Next
//	}
//	if numberOfNodes == 2 && n == 1 {
//		head.Next = nil
//	}
//
//	return head
//}
