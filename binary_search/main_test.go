package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  []int
		target int
		output int
	}{
		{
			input:  []int{1, 2, 3, 1},
			target: 2,
			output: 1,
		},
	}

	for _, tt := range testCases {
		got := search(tt.input, tt.target)
		assert.Equal(t, tt.output, got)
	}
}
