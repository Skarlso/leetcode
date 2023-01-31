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
			input:  []int{1, 1, 2, 3, 3},
			output: []int{1, 3},
		},
	}

	for _, tt := range testCases {
		got := findDuplicates(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
