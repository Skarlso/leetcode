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
			input:  []int{1, 1, 2},
			output: [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}},
		},
	}

	for _, tt := range testCases {
		got := permute(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
