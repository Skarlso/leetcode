package binarytreepath

import (
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	paths := binaryTreePaths(root)

	fmt.Println(paths)
}
