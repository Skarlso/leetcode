package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{2, 2, 1},
			output: 1,
		},
		{
			input:  []int{4, 1, 2, 1, 2},
			output: 4,
		},
		{
			input:  []int{1},
			output: 1,
		},
	}

	for _, tt := range testCases {
		got := singleNumber(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
