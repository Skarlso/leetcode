package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{3, 0, 1},
			output: 2,
		},
		{
			input:  []int{0, 1},
			output: 2,
		},
		{
			input:  []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			output: 8,
		},
	}

	for _, tt := range testCases {
		got := missingNumber(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
