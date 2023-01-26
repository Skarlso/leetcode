package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{3, 2, 3},
			output: 3,
		},
		{
			input:  []int{3, 3, 4},
			output: 3,
		},
		{
			input:  []int{-1, 1, 1, 1, 2, 1},
			output: 1,
		},
	}

	for _, tt := range testCases {
		got := majorityElement(tt.input)
		assert.Equal(t, tt.output, got)
	}
}

func TestLeetCodeV3(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{3, 2, 3},
			output: 3,
		},
		{
			input:  []int{3, 3, 4},
			output: 3,
		},
		{
			input:  []int{-1, 1, 1, 1, 2, 1},
			output: 1,
		},
		{
			input:  []int{1},
			output: 1,
		},
	}

	for _, tt := range testCases {
		got := majorityElementV3(tt.input)
		assert.Equal(t, tt.output, got)
	}
}

func TestLeetCodeV2(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{3, 2, 3},
			output: 3,
		},
		{
			input:  []int{3, 3, 4},
			output: 3,
		},
		{
			input:  []int{-1, 1, 1, 1, 2, 1},
			output: 1,
		},
		{
			input:  []int{1},
			output: 1,
		},
	}

	for _, tt := range testCases {
		got := majorityElementV2(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
