package main

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	n := addTwoNumbers(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	})

	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}
