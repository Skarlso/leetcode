package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  [][]int
		output [][]int
	}{
		{
			input:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			output: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
	}

	for _, tt := range testCases {
		rotate(tt.input)
		assert.Equal(t, tt.output, tt.input)
	}
}
