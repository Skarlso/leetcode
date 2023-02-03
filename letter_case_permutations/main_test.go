package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  string
		output []string
	}{
		{
			input:  "a1b2",
			output: []string{"a1b2", "A1b2", "a1B2", "A1B2"},
		},
	}

	for _, tt := range testCases {
		got := letterCasePermutation(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
