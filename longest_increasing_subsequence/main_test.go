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
			input:  []int{10, 9, 2, 5, 3, 7, 101, 18},
			output: 4,
		},
	}

	for _, tt := range testCases {
		got := lengthOfLIS(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
