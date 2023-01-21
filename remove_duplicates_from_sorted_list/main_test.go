package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  func() *ListNode
		output *ListNode
	}{
		{
			input: func() *ListNode {
				start := &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 3,
						},
					},
				}

				return start
			},
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
		{
			input: func() *ListNode {
				start := &ListNode{
					Val: 7,
					Next: &ListNode{
						Val: 7,
						Next: &ListNode{
							Val: 7,
							Next: &ListNode{
								Val: 7,
							},
						},
					},
				}

				return start
			},
			output: &ListNode{
				Val: 7,
			},
		},
		{
			input: func() *ListNode {
				start := &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: 0,
							Next: &ListNode{
								Val: 0,
							},
						},
					},
				}

				return start
			},
			output: &ListNode{
				Val: 0,
			},
		},
	}

	for _, tt := range testCases {
		got := deleteDuplicates(tt.input())
		assert.Equal(t, tt.output, got)
	}
}
