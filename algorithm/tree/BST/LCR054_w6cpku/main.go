package main

/* 把二叉搜索树转换为累加树 */

// 和538题一样
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
	sum := 0
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Right)
		sum += root.Val
		root.Val = sum
		traverse(root.Left)
	}
	traverse(root)
	return root
}
