package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		candidates []int
		target     int
		output     int
	}{
		{
			candidates: []int{1, 1, 1, 1, 1},
			target:     3,
			output:     5,
		},
	}

	for _, tt := range testCases {
		got := findTargetSumWays(tt.candidates, tt.target)
		assert.Equal(t, tt.output, got)
	}
}
