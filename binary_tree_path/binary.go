package binarytreepath

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	allPaths := []string{}

	dfs(root, &allPaths, nil)

	return allPaths
}

func dfs(node *TreeNode, allPaths *[]string, currentPath []string) {
	currentPath = append(currentPath, strconv.Itoa(node.Val))

	if node.Left == nil && node.Right == nil {
		*allPaths = append(*allPaths, strings.Join(currentPath, "->"))

		return
	}

	if node.Left != nil {
		dfs(node.Left, allPaths, currentPath)
	}

	if node.Right != nil {
		dfs(node.Right, allPaths, currentPath)
	}
}
