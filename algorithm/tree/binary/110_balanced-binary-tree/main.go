package main

import "math"

/* 平衡二叉树 */

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var traverse func(root *TreeNode) int
	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := traverse(root.Left)
		right := traverse(root.Right)
		if left == -1 || right == -1 || int(math.Abs(float64(left-right))) > 1 {
			return -1
		}
		return int(math.Max(float64(left), float64(right))) + 1
	}
	return traverse(root) >= 0
}
