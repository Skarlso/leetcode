package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{7, 1, 5, 3, 6, 4},
			output: 5,
		},
		{
			input:  []int{7, 6, 4, 3, 1},
			output: 0,
		},
		{
			input:  []int{1, 2},
			output: 1,
		},
		{
			input:  []int{2, 1, 2, 1, 0, 1, 2},
			output: 2,
		},
	}

	for _, tt := range testCases {
		got := maxProfit(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
