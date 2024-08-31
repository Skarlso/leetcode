package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)
	assert.Equal(t, 5, k)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 0, 2, 1, 3, 1}, nums)
}

func TestMain2(t *testing.T) {
	nums := []int{1, 1, 2}
	k := removeDuplicates(nums)
	assert.Equal(t, 2, k)
	assert.Equal(t, []int{1, 2, 1}, nums)
}
