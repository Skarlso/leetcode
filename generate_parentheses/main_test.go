package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		n      int
		output []string
	}{
		{
			n:      3,
			output: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}

	for _, tt := range testCases {
		got := generateParenthesis(tt.n)
		assert.Equal(t, tt.output, got)
	}
}
