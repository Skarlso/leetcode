package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  []int
		output bool
	}{
		{
			input:  []int{1, 2, 3, 1},
			output: true,
		},
	}

	for _, tt := range testCases {
		got := containsDuplicate(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
