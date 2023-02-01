package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  [][]byte
		target string
		output bool
	}{
		{
			input:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			target: "ABCCED",
			output: true,
		},
		{
			input:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			target: "SEE",
			output: true,
		},
		{
			input:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			target: "ABCB",
			output: false,
		},
	}

	for _, tt := range testCases {
		got := exist(tt.input, tt.target)
		assert.Equal(t, tt.output, got)
	}
}
