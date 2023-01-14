package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		init        []int
		left, right int
		output      int
	}{
		{
			init:   []int{-2, 0, 3, -5, 2, -1},
			left:   0,
			right:  2,
			output: 1,
		},
		{
			init:   []int{-2, 0, 3, -5, 2, -1},
			left:   2,
			right:  5,
			output: -1,
		},
		{
			init:   []int{-2, 0, 3, -5, 2, -1},
			left:   0,
			right:  5,
			output: -3,
		},
		{
			init:   []int{-1},
			left:   0,
			right:  0,
			output: -1,
		},
	}

	for _, tt := range testCases {
		sum := Constructor(tt.init)
		got := sum.SumRange(tt.left, tt.right)
		assert.Equal(t, tt.output, got)
	}
}
