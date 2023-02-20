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
		// {
		// 	input:  []int{2, 3, -2, 4},
		// 	output: 6,
		// },
		{
			input:  []int{-2, 3, -4},
			output: 24,
		},
	}

	for _, tt := range testCases {
		got := maxProduct(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
