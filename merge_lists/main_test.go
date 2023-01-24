package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input1 func() *ListNode
		input2 func() *ListNode
		output *ListNode
	}{
		{
			input1: func() *ListNode {
				start := &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
						},
					},
				}

				return start
			},
			input2: func() *ListNode {
				start := &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				}

				return start
			},
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		got := mergeTwoLists(tt.input1(), tt.input2())
		assert.Equal(t, tt.output, got)
	}
}

func TestLeetCodeV2(t *testing.T) {
	testCases := []struct {
		input1 func() *ListNode
		input2 func() *ListNode
		output *ListNode
	}{
		{
			input1: func() *ListNode {
				start := &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
						},
					},
				}

				return start
			},
			input2: func() *ListNode {
				start := &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				}

				return start
			},
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		got := mergeTwoListsRecursive(tt.input1(), tt.input2())
		assert.Equal(t, tt.output, got)
	}
}
