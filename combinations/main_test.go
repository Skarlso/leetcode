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
			n:      4,
			k:      2,
			output: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
	}

	for _, tt := range testCases {
		got := combine(tt.n, tt.k)
		assert.Equal(t, tt.output, got)
	}
}
