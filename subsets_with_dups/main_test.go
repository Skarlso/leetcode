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
			input:  []int{1, 2, 2},
			output: [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
	}

	for _, tt := range testCases {
		got := subsetsWithDup(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
