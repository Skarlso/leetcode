package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		input  func() *ListNode
		output *ListNode
	}{
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
			output: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}

	for _, tt := range testCases {
		got := reverseList(tt.input())
		assert.Equal(t, tt.output, got)
	}
}
