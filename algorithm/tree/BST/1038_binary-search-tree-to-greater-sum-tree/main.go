package main

/* 从二叉搜索树到更大和树 */

// 和538一样
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

func bstToGst(root *TreeNode) *TreeNode {
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
