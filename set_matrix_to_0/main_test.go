package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  [][]int
		output [][]int
	}{
		{
			input:  [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			output: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
	}
	for _, tt := range testCases {
		setZeroes(tt.input)
		assert.Equal(t, tt.output, tt.input)
	}
}

func TestLeetCodeConstant(t *testing.T) {
	testCases := []struct {
		input  [][]int
		output [][]int
	}{
		{
			input:  [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			output: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			input:  [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			output: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}
	for _, tt := range testCases {
		setZeroConstantMemory(tt.input)
		assert.Equal(t, tt.output, tt.input)
	}
}

func BenchmarkNotConstantMemory(b *testing.B) {
	input := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	// output := [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}

	for i := 0; i < b.N; i++ {
		setZeroes(input)
	}
}

func BenchmarkConstantMemory(b *testing.B) {
	input := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	// output := [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}

	for i := 0; i < b.N; i++ {
		setZeroConstantMemory(input)
	}
}
