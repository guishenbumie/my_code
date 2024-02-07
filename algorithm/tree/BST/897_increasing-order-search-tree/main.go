package main

/* 递增顺序搜索树 */

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

func increasingBST(root *TreeNode) *TreeNode {
	dummy := &TreeNode{}
	ans := dummy
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)

		ans.Right = root
		root.Left = nil
		ans = root

		traverse(root.Right)
	}
	traverse(root)
	return dummy.Right
}
