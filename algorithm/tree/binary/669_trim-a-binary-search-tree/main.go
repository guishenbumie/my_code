package main

/* 修剪二叉搜索树 */

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

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	//如果当前根节点比边界小，根以及它的左孩子都不符合条件，都要舍弃，他的右孩子变为根节点
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	//如果当前根节点比边界大，根以及它的右孩子都不符合条件，都要舍弃，他的左孩子变为根节点
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
