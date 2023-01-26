package main

// Definition for singly-linked list.
type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

// The two lists are sorted. So just go an merge them. I hate switching out nodes.
// I always confuse how to fuck go to the next node.
// func mergeTrees(list1 *TreeNode, list2 *TreeNode) *TreeNode {
// 	if list1 == nil {
// 		return list2
// 	}
// 	if list2 == nil {
// 		return list1
// 	}

// 	newList := list1

// 	if list1.Val > list2.Val {
// 		newList = list2
// 		list2 = list2.Next
// 	} else {
// 		// go to the second item in the list.
// 		list1 = list1.Next
// 	}

// 	// Create a new list value so we don't loose the begin of the original new list.
// 	current := newList

// 	for list1 != nil && list2 != nil {
// 		if list1.Val < list2.Val {
// 			current.Next = list1
// 			list1 = list1.Next
// 		} else {
// 			current.Next = list2
// 			list2 = list2.Next
// 		}
// 		// Move to the next item. Otherwise the current item would keep getting overwritten.
// 		current = current.Next
// 	}
// 	// Add the remaining items from either list1 or list2.
// 	if list1 != nil {
// 		current.Next = list1
// 	} else if list2 != nil {
// 		current.Next = list2
// 	}
// 	return newList
// }

func mergeTrees(list1 *TreeNode, list2 *TreeNode) *TreeNode {
	if list1 != nil && list2 != nil {
		root := &TreeNode{
			Val:   list1.Val + list2.Val,
			Right: mergeTrees(list1.Right, list2.Right),
			Left:  mergeTrees(list1.Left, list2.Left),
		}
		return root
	} else if list1 != nil && list2 == nil {
		return list1
	} else if list1 == nil && list2 != nil {
		return list2
	}
	return nil
}
