package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingBits(t *testing.T) {
	testCases := []struct {
		input  int
		output []int
	}{
		{
			input:  2,
			output: []int{0, 1, 1},
		},
		{
			input:  5,
			output: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, tt := range testCases {
		got := countBits(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
