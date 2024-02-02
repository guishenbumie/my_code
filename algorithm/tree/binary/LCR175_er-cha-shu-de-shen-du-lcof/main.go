package main

import "math"

/* 计算二叉树的深度 */

// 和104题一样
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

func calculateDepth(root *TreeNode) int {
	var traverse func(root *TreeNode) int
	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftMax := traverse(root.Left)
		rightMax := traverse(root.Right)
		max := int(math.Max(float64(leftMax), float64(rightMax)))
		return max + 1
	}
	return traverse(root)
}
