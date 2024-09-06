package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
	m := make(map[int]struct{}, 0)
	for _, n := range nums {
		m[n] = struct{}{}
	}

	dummy := &ListNode{Val: -1, Next: head}
	prev := dummy
	for prev.Next != nil {
		if _, ok := m[prev.Next.Val]; ok {
			// don't shift prev
			prev.Next = prev.Next.Next

			continue
		} else {
			prev = prev.Next
		}
	}

	return dummy.Next
}
