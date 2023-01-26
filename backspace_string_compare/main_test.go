package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCode(t *testing.T) {
	testCases := []struct {
		s      string
		t      string
		output bool
	}{
		{
			s:      "ab#c",
			t:      "ad#c",
			output: true,
		},
		{
			s:      "ab##",
			t:      "c#d#",
			output: true,
		},
		{
			s:      "a#c",
			t:      "b",
			output: false,
		},
		{
			s:      "y#fo##f",
			t:      "y#f#o##f",
			output: true,
		},
	}

	for _, tt := range testCases {
		got := backspaceCompare(tt.s, tt.t)
		assert.Equal(t, tt.output, got)
	}
}

func BenchmarkCompare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		got := backspaceCompare("y#f#o##f", "y#f#o##f")
		if !got {
			b.Fatal()
		}
	}
}
func BenchmarkCompareString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		got := backspaceCompareString("y#f#o##f", "y#f#o##f")
		if !got {
			b.Fatal()
		}
	}
}
