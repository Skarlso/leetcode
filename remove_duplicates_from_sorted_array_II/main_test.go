package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := removeDuplicates(nums)
	assert.Equal(t, 5, k)
	assert.Equal(t, []int{1, 1, 2, 2, 3, 3}, nums)
}

func TestMain2(t *testing.T) {
	nums := []int{1, 1, 2}
	k := removeDuplicates(nums)
	assert.Equal(t, 2, k)
	assert.Equal(t, []int{1, 2, 1}, nums)
}
