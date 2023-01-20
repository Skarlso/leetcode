package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		input  *BinaryTree
		output []int
	}{
		{
			input: &BinaryTree{
				Value: 1,
				Right: &BinaryTree{
					Value: 2,
					Right: &BinaryTree{
						Value: 3,
					},
					Left: &BinaryTree{
						Value: 4,
					},
				},
				Left: &BinaryTree{
					Value: 5,
				},
			},
			output: []int{6, 7, 6},
		},
	}

	for _, tt := range testCases {
		got := BranchSums(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
