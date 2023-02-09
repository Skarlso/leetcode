package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		n      int
		k      int
		output [][]int
	}{
		{
			n:      7,
			k:      3,
			output: [][]int{{1, 2, 4}},
		},
	}

	for _, tt := range testCases {
		got := combinationSum3(tt.k, tt.n)
		assert.Equal(t, tt.output, got)
	}
}
