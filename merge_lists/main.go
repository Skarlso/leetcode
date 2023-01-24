package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// The two lists are sorted. So just go an merge them. I hate switching out nodes.
// I always confuse how to fuck go to the next node.
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	newList := list1

	if list1.Val > list2.Val {
		newList = list2
		list2 = list2.Next
	} else {
		// go to the second item in the list.
		list1 = list1.Next
	}

	// Create a new list value so we don't loose the begin of the original new list.
	current := newList

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		// Move to the next item. Otherwise the current item would keep getting overwritten.
		current = current.Next
	}
	// Add the remaining items from either list1 or list2.
	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}
	return newList
}

func mergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoListsRecursive(list1.Next, list2)
		return list1
	}

	list2.Next = mergeTwoListsRecursive(list1, list2.Next)
	return list2
}
