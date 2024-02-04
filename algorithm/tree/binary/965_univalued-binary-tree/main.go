package main

/* 单值二叉树 */

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

func isUnivalTree(root *TreeNode) bool {
	val := root.Val
	var traverse func(root *TreeNode) bool
	traverse = func(root *TreeNode) bool {
		if root == nil {
			return true
		}
		if root.Val != val {
			return false
		}
		if !traverse(root.Left) {
			return false
		}
		return traverse(root.Right)
	}
	return traverse(root)
}
