package main

/* 把二叉搜索树转换为累加树 */

// 和1038一样
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

func convertBST(root *TreeNode) *TreeNode {
	var sum int
	var f func(r *TreeNode)
	f = func(r *TreeNode) {
		if r == nil {
			return
		}
		f(r.Right)
		sum += r.Val
		r.Val = sum
		f(r.Left)
	}
	f(root)
	return root
}
