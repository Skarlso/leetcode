package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  *TreeNode
		output int
	}{
		{
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 4,
					},
					Left: &TreeNode{
						Val: 5,
					},
				},
				Left: &TreeNode{
					Val: 3,
				},
			},
			output: 3,
		},
		{
			input: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			output: 1,
		},
	}

	for _, tt := range testCases {
		got := diameterOfBinaryTree(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
