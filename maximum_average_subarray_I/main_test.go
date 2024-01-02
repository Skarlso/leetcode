package main

import (
	"testing"
)

func TestMaximum(t *testing.T) {
	result := findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)
	if result != 12.75 {
		t.Fatal("expected 12.75, got: ", result)
	}

	result = findMaxAverage([]int{-1}, 1)
	if result != -1 {
		t.Fatal("expected -1, got: ", result)
	}
}
