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
			input:  []int{2, 3, 2},
			output: 3,
		},
	}

	for _, tt := range testCases {
		got := rob(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
