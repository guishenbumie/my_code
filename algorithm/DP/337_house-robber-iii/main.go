package main

import "math"

/* 打家劫舍 III */

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

func rob(root *TreeNode) int {
	memo := make(map[*TreeNode]int)
	var dp func(node *TreeNode) int
	dp = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if ans, ok := memo[node]; ok {
			return ans
		}
		//偷当前点，左右子节点不偷，左右子节点的孩子偷
		do := node.Val
		if node.Left != nil {
			do += dp(node.Left.Left) + dp(node.Left.Right)
		}
		if node.Right != nil {
			do += dp(node.Right.Left) + dp(node.Right.Right)
		}
		//不偷当前点，左右子节点偷
		notDo := dp(node.Left) + dp(node.Right)
		memo[node] = int(math.Max(
			float64(do),
			float64(notDo),
		))
		return memo[node]
	}
	return dp(root)
}
