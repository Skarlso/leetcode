package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsCycle(t *testing.T) {
	testCases := []struct {
		input  func() *ListNode
		output bool
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
				next.Next = start
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
	}

	for _, tt := range testCases {
		got := hasCycle(tt.input())
		assert.Equal(t, tt.output, got)
	}
}
