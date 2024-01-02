package main

/* 二叉搜索树中的搜索 */

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

// 普通二叉树的搜索方式，但BST的特点，左小右大，可以有更好的搜索方式
//func searchBST(root *TreeNode, val int) *TreeNode {
//	if root == nil {
//		return nil
//	}
//	if root.Val == val {
//		return root
//	}
//	left := searchBST(root.Left, val)
//	right := searchBST(root.Right, val)
//	if left != nil {
//		return left
//	} else {
//		return right
//	}
//}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	//去左子树搜
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	//去右子树搜
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	return root
}
