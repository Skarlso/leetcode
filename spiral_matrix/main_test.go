package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  [][]int
		output []int
	}{
		{
			input:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			output: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
	}

	for _, tt := range testCases {
		got := spiralOrder(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
