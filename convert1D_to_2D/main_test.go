package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  []int
		m      int
		n      int
		output [][]int
	}{
		{
			input: []int{1, 2, 3, 4},
			m:     2,
			n:     2,
			output: [][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			input:  []int{1, 2},
			m:      1,
			n:      1,
			output: [][]int{},
		},
		{
			input:  []int{5},
			m:      3,
			n:      1,
			output: [][]int{},
		},
		{
			input:  []int{6, 2, 6, 8},
			m:      4,
			n:      2,
			output: [][]int{},
		},
		{
			input: []int{1, 1, 1, 1},
			m:     4,
			n:     1,
			output: [][]int{
				{1},
				{1},
				{1},
				{1},
			},
		},
		{
			input:  []int{1, 2, 3},
			m:      1,
			n:      3,
			output: [][]int{{1, 2, 3}},
		},
		// {
		// 	input:  []int{59, 48, 86, 96, 90, 70, 54, 76, 96, 94, 20, 54, 84, 13, 63, 79, 53, 36, 5, 89, 48, 28, 57, 57, 31, 47, 13, 88, 48, 70, 83, 46, 8, 51, 36, 25, 16, 59, 12, 85, 28, 83, 82, 77, 3, 7, 43, 32, 17, 73, 27, 9, 98, 71, 27, 76, 28, 30, 43, 100, 19, 77, 1, 36, 19, 17, 70, 33, 23, 15, 85, 4, 17, 2, 46, 31, 76, 42, 3, 23, 34, 7, 17, 95, 54, 68, 15, 96, 9, 61, 58, 100, 47, 66, 76, 14, 58, 65, 29, 79, 6, 94, 85, 67, 39, 59, 83, 4, 8, 62, 52, 55, 88, 14, 63, 92, 57, 10, 11, 42, 78, 19, 9, 17, 36, 80, 82, 75, 71, 92, 42, 49, 77, 68, 38, 59, 65, 78, 56, 13, 90, 15, 84, 84, 99, 71, 93, 28, 56, 58, 30, 53, 44, 59, 52, 54, 94, 13, 41, 55, 41, 86, 33, 48, 20, 88, 61, 9, 12, 11, 12, 56, 13, 96, 95, 48, 17, 72, 8, 76, 72, 72, 43, 86, 16, 53, 46, 86, 97, 55, 93, 7, 37, 77, 40, 46, 14, 74, 65},
		// 	m:      8,
		// 	n:      18,
		// 	output: [][]int{{1, 2, 3}},
		// },
	}

	for _, tt := range testCases {
		got := construct2DArray(tt.input, tt.m, tt.n)
		assert.Equal(t, tt.output, got)
	}
}