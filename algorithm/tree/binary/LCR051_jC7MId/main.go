package main

import "math"

/* 二叉树中的最大路径和 */

// 和124题一样
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

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftMax := max(dfs(root.Left), 0)
		rightMax := max(dfs(root.Right), 0)
		ans = max(ans, leftMax+rightMax+root.Val)
		return max(leftMax, rightMax) + root.Val
	}
	dfs(root)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
