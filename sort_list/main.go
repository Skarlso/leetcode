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
func sortListBubbleV1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	fast := head.Next
	slow := head

	swapped := false
	for {

		if slow.Val > fast.Val {
			slow.Val, fast.Val = fast.Val, slow.Val

			swapped = true
		}

		slow = slow.Next
		fast = fast.Next

		// reset if we reached the end
		if fast == nil {
			fast = head.Next
			slow = head
			if !swapped {
				break
			}

			// start a new round
			swapped = false
		}
	}

	return head
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// split the list
	left := head
	right := getMid(head)
	right, right.Next = right.Next, nil

	// sort
	left = sortList(left)
	right = sortList(right)

	// merge
	return merge(left, right)
}

func getMid(node *ListNode) *ListNode {
	slow, fast := node, node.Next

	// fast jumps by two and by the time it gets to the end, slow will be in the middle
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for left != nil && right != nil {
		if left.Val < right.Val {
			tail.Next = left
			left = left.Next
		} else {
			tail.Next = right
			right = right.Next
		}

		tail = tail.Next
	}

	if left != nil {
		tail.Next = left
	}

	if right != nil {
		tail.Next = right
	}

	return dummy.Next
}
