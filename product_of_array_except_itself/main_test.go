package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{-1, 1, 0, -3, 3},
			output: []int{0, 0, 9, 0, 0},
		},
		{
			input:  []int{0, 0},
			output: []int{0, 0},
		},
	}

	for _, tt := range testCases {
		got := productExceptSelf(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
