package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input1 *TreeNode
		input2 *TreeNode
		output bool
	}{
		{
			input1: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			input2: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			output: true,
		},
	}

	for _, tt := range testCases {
		got := isSubtree(tt.input1, tt.input2)
		assert.Equal(t, tt.output, got)
	}
}
