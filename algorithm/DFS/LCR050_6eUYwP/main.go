package main

/* 路径总和 III */

// 和437题一样
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

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var traverse func(root *TreeNode, targetSum int) int
	traverse = func(root *TreeNode, targetSum int) int {
		res := 0
		if root == nil {
			return res
		}
		val := root.Val
		if val == targetSum {
			res++
		}
		res += traverse(root.Left, targetSum-val)
		res += traverse(root.Right, targetSum-val)
		return res
	}
	ans := traverse(root, targetSum)
	ans += pathSum(root.Left, targetSum)
	ans += pathSum(root.Right, targetSum)
	return ans
}
