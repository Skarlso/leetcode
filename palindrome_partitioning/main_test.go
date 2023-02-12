package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  string
		output [][]string
	}{
		{
			input:  "aab",
			output: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
	}

	for _, tt := range testCases {
		got := partition(tt.input)
		assert.Equal(t, tt.output, got)
	}
}

func TestLeetCodeIsPalindrome(t *testing.T) {
	testCases := []struct {
		input  string
		output bool
	}{
		{
			input:  "aa",
			output: true,
		},
		{
			input:  "aba",
			output: true,
		},
		{
			input:  "axca",
			output: false,
		},
	}

	for _, tt := range testCases {
		got := isPalindrome(tt.input, 0, len(tt.input)-1)
		assert.Equal(t, tt.output, got)
	}
}
