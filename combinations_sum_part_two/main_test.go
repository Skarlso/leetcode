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
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			output:     [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}},
		},
	}

	for _, tt := range testCases {
		got := combinationSum2(tt.candidates, tt.target)
		assert.Equal(t, tt.output, got)
	}
}
