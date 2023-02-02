package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{100, 4, 200, 1, 3, 2},
			output: 4,
		},
		{
			input:  []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			output: 9,
		},
	}

	for _, tt := range testCases {
		got := longestConsecutive(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
