package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		input  []int
		output bool
	}{
		{
			input:  []int{1, 2, 3, 1},
			output: true,
		},
		{
			input:  []int{1, 2, 3, 4},
			output: false,
		},
		{
			input:  []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			output: true,
		},
	}

	for _, tt := range testCases {
		got := containsDuplicate(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
