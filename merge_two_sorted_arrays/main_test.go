package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	m1 := []int{1, 2, 3, 0, 0, 0}
	m2 := []int{2, 4, 5}
	merge(m1, 3, m2, len(m2))

	assert.Equal(t, []int{1, 2, 2, 3, 4, 5}, m1)

	m1 = []int{1}
	m2 = []int{}
	merge(m1, 1, m2, len(m2))

	assert.Equal(t, []int{1}, m1)

	m1 = []int{0}
	m2 = []int{1}
	merge(m1, 0, m2, len(m2))

	assert.Equal(t, []int{1}, m1)

	m1 = []int{2, 0}
	m2 = []int{1}
	merge(m1, 1, m2, len(m2))

	assert.Equal(t, []int{1, 2}, m1)

	m1 = []int{4, 0, 0, 0, 0, 0}
	m2 = []int{1, 2, 3, 5, 6}
	merge(m1, 1, m2, len(m2))

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, m1)
}
