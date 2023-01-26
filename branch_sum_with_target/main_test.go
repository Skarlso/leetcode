package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		input  *TreeNode
		target int
		output bool
	}{
		{
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
					Left: &TreeNode{
						Val: 4,
					},
				},
				Left: &TreeNode{
					Val: 5,
				},
			},
			target: 6,
			output: true,
		},
	}

	for _, tt := range testCases {
		got := hasPathSum(tt.input, tt.target)
		assert.Equal(t, tt.output, got)
	}
}

func TestContainsDuplicateRec(t *testing.T) {
	testCases := []struct {
		input  *TreeNode
		target int
		output bool
	}{
		{
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
					Left: &TreeNode{
						Val: 4,
					},
				},
				Left: &TreeNode{
					Val: 5,
				},
			},
			target: 6,
			output: true,
		},
		{
			input: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 8,
					Right: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 1,
						},
					},
					Left: &TreeNode{
						Val: 13,
					},
				},
				Left: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 11,
						Left: &TreeNode{
							Val: 7,
						},
						Right: &TreeNode{
							Val: 2,
						},
					},
				},
			},
			target: 22,
			output: true,
		},
		{
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			target: 1,
			output: false,
		},
	}

	for _, tt := range testCases {
		got := hasPathSumRec(tt.input, tt.target)
		assert.Equal(t, tt.output, got)
	}
}
