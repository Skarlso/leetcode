package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  []int
		output [][]int
	}{
		{
			input:  []int{1, 2, 3},
			output: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}

	for _, tt := range testCases {
		got := subsets(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
