package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// The trivial solution would be to get it out into a list then performs a palindrome check on it.
// But that's not what we aim for.
// If we loop twice that's still O(n) technically.
// We could determine the middle point and read up until then.
// Then read the rest. OH! Fast, Slow pointer. We get the middle one.
// Gather the values from the Slow pointer... Gather the values from the fast pointer.
// As the slow will be in the middle, we just continue from Slow until we get to the end.
// then compare the values.
// Reverse and continue?
// O(1) space!
// Reverse the _SECOND HALF_ of the list instead of the first half.
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next

		// REVERSE HALF OF THE LIST
		// Slow will have the rest
		// Prev will have the pointer to the previous items.
		// next = slow.Next
		// slow.Next = prev
		// prev = slow
		// slow = next
		slow = slow.Next
	}

	// Reverse second half
	var prev *ListNode

	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	for prev != nil {
		if head.Val != prev.Val {
			return false
		}

		head = head.Next
		prev = prev.Next
	}

	// update the final missing step
	// prev = slow

	// Now read prev and slow simultaneously and compare the values. At the point when BOTH are nil, return true
	// if either one of them is nil before running out of values, break and return false.

	// for {
	// 	if slow == nil && prev != nil || prev == nil && slow != nil {
	// 		break
	// 	}

	// 	if slow == nil && prev == nil {
	// 		return true
	// 	}

	// 	if slow.Val != prev.Val {
	// 		return false
	// 	}

	// 	slow = slow.Next
	// 	prev = prev.Next
	// }

	return true
}
