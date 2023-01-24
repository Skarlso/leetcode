package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		input  []byte
		target byte
		output byte
	}{
		{
			input:  []byte{'c', 'f', 'j'},
			target: 'c',
			output: 'f',
		},
		{
			input:  []byte{'x', 'x', 'y', 'y'},
			target: 'z',
			output: 'x',
		},
		{
			input:  []byte{'c', 'f', 'j'},
			target: 'g',
			output: 'j',
		},
	}

	for _, tt := range testCases {
		got := nextGreatestLetter(tt.input, tt.target)
		assert.Equal(t, tt.output, got, fmt.Sprintf("got: %s", string(got)))
	}
}

func TestLeetCodeV2(t *testing.T) {
	testCases := []struct {
		input  []byte
		target byte
		output byte
	}{
		{
			input:  []byte{'c', 'f', 'j'},
			target: 'c',
			output: 'f',
		},
		{
			input:  []byte{'x', 'x', 'y', 'y'},
			target: 'z',
			output: 'x',
		},
		{
			input:  []byte{'c', 'f', 'j'},
			target: 'g',
			output: 'j',
		},
	}

	for _, tt := range testCases {
		got := nextGreatestLetterV2(tt.input, tt.target)
		assert.Equal(t, tt.output, got, fmt.Sprintf("got: %s", string(got)))
	}
}
