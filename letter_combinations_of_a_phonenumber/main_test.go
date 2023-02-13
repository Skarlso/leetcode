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
			input:  "23",
			output: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}

	for _, tt := range testCases {
		got := letterCombinations(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
