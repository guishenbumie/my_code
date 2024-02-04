package main

import "math"

/* 二叉搜索树节点最小距离 */

// 和530题一样
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

func minDiffInBST(root *TreeNode) int {
	ans := math.MaxInt32
	var prev *TreeNode
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		if prev != nil {
			ans = int(math.Min(float64(ans), float64(root.Val-prev.Val)))
		}
		prev = root
		traverse(root.Right)
	}
	traverse(root)
	return ans
}
