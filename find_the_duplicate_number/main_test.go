package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{-1, 1, 0, -3, 3},
			output: []int{0, 0, 9, 0, 0},
		},
		{
			input:  []int{0, 0},
			output: []int{0, 0},
		},
	}

	for _, tt := range testCases {
		got := productExceptSelf(tt.input)
		assert.Equal(t, tt.output, got)
	}
}

func TestLeetCodeV2(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 3, 4, 2, 2},
			output: 2,
		},
		{
			input:  []int{1, 3, 2, 4, 2},
			output: 2,
		},
	}

	for _, tt := range testCases {
		got := findDuplicate(tt.input)
		assert.Equal(t, tt.output, got)
	}
}

func TestLeetCodeV3(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 3, 4, 2, 2},
			output: 2,
		},
		{
			input:  []int{1, 3, 2, 4, 2},
			output: 2,
		},
		{
			input:  []int{3, 2, 1, 5, 3, 4},
			output: 3,
		},
	}

	for _, tt := range testCases {
		got := findDuplicateV3(tt.input)
		assert.Equal(t, tt.output, got)
	}
}

func TestLeetCodeV4(t *testing.T) {
	testCases := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 3, 4, 2, 2},
			output: 2,
		},
		{
			input:  []int{1, 3, 2, 4, 2},
			output: 2,
		},
		{
			input:  []int{3, 2, 1, 5, 3, 4},
			output: 3,
		},
	}

	for _, tt := range testCases {
		got := findDuplicateV4(tt.input)
		assert.Equal(t, tt.output, got)
	}
}
