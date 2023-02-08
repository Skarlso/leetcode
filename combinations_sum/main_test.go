package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		candidates []int
		target     int
		output     [][]int
	}{
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			output:     [][]int{{2, 2, 3}, {7}},
		},
	}

	for _, tt := range testCases {
		got := combinationSum(tt.candidates, tt.target)
		assert.Equal(t, tt.output, got)
	}
}
