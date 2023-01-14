package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbingStairs(t *testing.T) {
	testCases := []struct {
		input  int
		output int
	}{
		{
			input:  2,
			output: 2,
		},
		{
			input:  3,
			output: 3,
		},
	}

	for _, tt := range testCases {
		got := climbStairs(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
