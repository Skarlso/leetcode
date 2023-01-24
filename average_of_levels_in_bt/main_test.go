package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  *TreeNode
		output []float64
	}{
		{
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			output: []float64{3.00000, 14.50000, 11.00000},
		},
	}

	for _, tt := range testCases {
		got := averageOfLevels(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
