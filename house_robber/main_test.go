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
			input:  []int{1, 2, 3, 1},
			output: 4,
		},
	}

	for _, tt := range testCases {
		got := rob(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
