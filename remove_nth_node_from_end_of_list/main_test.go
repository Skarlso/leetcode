package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElements(t *testing.T) {
	testCases := []struct {
		input  func() *ListNode
		value  int
		output *ListNode
	}{
		{
			input: func() *ListNode {
				return &ListNode{
					Val: 1,
				}
			},
			value:  1,
			output: nil,
		},
		{
			input: func() *ListNode {
				start := &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val: 6,
										Next: &ListNode{
											Val: 7,
										},
									},
								},
							},
						},
					},
				}

				return start
			},
			value: 6,
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
								Next: &ListNode{
									Val: 7,
								},
							},
						},
					},
				},
			},
		},
		{
			input: func() *ListNode {
				start := &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
					},
				}

				return start
			},
			value: 1,
			output: &ListNode{
				Val: 1,
			},
		},
		{
			input: func() *ListNode {
				start := &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
					},
				}

				return start
			},
			value: 2,
			output: &ListNode{
				Val: 2,
			},
		},
	}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("value-%d", tt.value), func(t *testing.T) {
			got := removeNthFromEnd(tt.input(), tt.value)
			assert.Equal(t, tt.output, got)
		})
	}
}
