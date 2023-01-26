package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input1 *TreeNode
		input2 *TreeNode
		output *TreeNode
	}{
		{
			input1: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			input2: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			output: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 7,
					},
				},
				Left: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 4,
					},
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
	}

	for _, tt := range testCases {
		got := mergeTrees(tt.input1, tt.input2)
		assert.Equal(t, tt.output, got)
	}
}
