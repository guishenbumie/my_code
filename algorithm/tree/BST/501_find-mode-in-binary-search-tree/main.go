package main

/* 二叉搜索树中的众数 */

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

func findMode(root *TreeNode) []int {
	ans := []int{}

	var base, count, maxCount int

	update := func(x int) {
		if x == base {
			count++
		} else {
			base = x
			count = 1
		}
		if count == maxCount {
			ans = append(ans, base)
		} else if count > maxCount {
			maxCount = count
			ans = []int{base}
		}
	}

	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		update(root.Val)
		traverse(root.Right)
	}
	traverse(root)
	return ans
}
