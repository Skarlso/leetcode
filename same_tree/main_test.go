package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeetCodeFunction(t *testing.T) {
	q := &TreeNode{
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
			},
		},
	}
	p := &TreeNode{
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
			},
		},
	}

	assert.True(t, isSameTree(q, p))
}

func TestLeetCodeFunctionRecursive(t *testing.T) {
	q := &TreeNode{
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
			},
		},
	}
	p := &TreeNode{
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
						Right: &TreeNode{
							Val: 5,
						},
					},
				},
			},
		},
	}

	assert.True(t, isSameTreeRec(q, p))
}

func TestLeetCodeFunctionRecursive2(t *testing.T) {
	q := &TreeNode{
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	p := &TreeNode{
		Right: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 5,
				},
				Left: &TreeNode{
					Val: 4,
				},
			},
			Left: &TreeNode{
				Val: 2,
			},
		},
	}

	assert.False(t, isSameTreeRec(q, p))
}
