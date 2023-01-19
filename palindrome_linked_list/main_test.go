package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input  func() *ListNode
		output bool
	}{
		{
			input: func() *ListNode {
				nextnextnext := &ListNode{
					Val: 1,
				}
				nextnext := &ListNode{
					Val:  2,
					Next: nextnextnext,
				}
				next := &ListNode{
					Val:  2,
					Next: nextnext,
				}
				start := &ListNode{
					Val:  1,
					Next: next,
				}

				return start
			},
			output: true,
		},
		{
			input: func() *ListNode {
				nextnext := &ListNode{
					Val: 1,
				}
				next := &ListNode{
					Val:  0,
					Next: nextnext,
				}
				start := &ListNode{
					Val:  1,
					Next: next,
				}

				return start
			},
			output: true,
		},
		{
			input: func() *ListNode {
				next := &ListNode{
					Val: 2,
				}
				start := &ListNode{
					Val:  1,
					Next: next,
				}

				return start
			},
			output: false,
		},
		{
			input: func() *ListNode {
				start := &ListNode{
					Val: 1,
				}

				return start
			},
			output: true,
		},
	}

	for _, tt := range testCases {
		got := isPalindrome(tt.input())
		assert.Equal(t, tt.output, got)
	}
}
