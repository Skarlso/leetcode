package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	input := []int{3, 2, 2, 3}
	got := removeElement(input, 3)
	assert.Equal(t, 2, got)
	assert.Equal(t, []int{2, 2, 3, 3}, input)
}

func TestRemoveElement2(t *testing.T) {
	input := []int{0, 1, 2, 2, 3, 0, 4, 2}
	got := removeElement(input, 2)
	assert.Equal(t, 5, got)
	assert.Equal(t, []int{0, 1, 4, 0, 3, 2, 2, 2}, input)
}

func TestRemoveElement3(t *testing.T) {
	input := []int{1}
	got := removeElement(input, 1)
	assert.Equal(t, 0, got)
	assert.Equal(t, []int{1}, input)
}

func TestRemoveElement4(t *testing.T) {
	input := []int{3, 3}
	got := removeElement(input, 3)
	assert.Equal(t, 0, got)
	assert.Equal(t, []int{3, 3}, input)
}
