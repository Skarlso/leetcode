package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllDisappearedNumbers(t *testing.T) {
	testCases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{4, 3, 2, 7, 8, 2, 3, 1},
			output: []int{5, 6},
		},
		{
			input:  []int{1, 1},
			output: []int{2},
		},
	}

	for _, tt := range testCases {
		got := findDisappearedNumbers(tt.input)
		assert.Equal(t, tt.output, got)
	}
}

func TestFindAllDisappearedNumbersV2(t *testing.T) {
	testCases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{4, 3, 2, 7, 8, 2, 3, 1},
			output: []int{5, 6},
		},
		{
			input:  []int{1, 1},
			output: []int{2},
		},
	}

	for _, tt := range testCases {
		got := findDisappearedNumbersV2(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
