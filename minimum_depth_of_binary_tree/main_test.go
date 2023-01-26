package main

import "testing"

func BenchmarkIterative(b *testing.B) {
	input := &TreeNode{
		Right: &TreeNode{
			Value: 1,
			Left: &TreeNode{
				Value: 2,
				Right: &TreeNode{
					Value: 3,
					Left: &TreeNode{
						Value: 4,
						Right: &TreeNode{
							Value: 5,
						},
					},
				},
			},
		},
	}

	for i := 0; i < b.N; i++ {
		got := minDepthIterative(input)
		if got != 6 {
			b.Fatalf("expected %d but got: %d", 5, got)
		}
	}
}

func BenchmarkRecursive(b *testing.B) {
	input := &TreeNode{
		Right: &TreeNode{
			Value: 1,
			Left: &TreeNode{
				Value: 2,
				Right: &TreeNode{
					Value: 3,
					Left: &TreeNode{
						Value: 4,
						Right: &TreeNode{
							Value: 5,
						},
					},
				},
			},
		},
	}

	for i := 0; i < b.N; i++ {
		got := minDepthRecursive(input)
		if got != 6 {
			b.Fatalf("expected %d but got: %d", 5, got)
		}
	}
}
