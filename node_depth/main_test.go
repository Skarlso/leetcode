package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		input  *BinaryTree
		output int
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
			output: 6,
		},
	}

	for _, tt := range testCases {
		got := NodeDepths(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
