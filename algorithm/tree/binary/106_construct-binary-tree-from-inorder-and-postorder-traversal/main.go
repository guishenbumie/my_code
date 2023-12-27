package main

/* 从中序与后序遍历序列构造二叉树 */

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

func buildTree(inorder []int, postorder []int) *TreeNode {
	return build(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func build(inorder []int, inStart, inEnd int, postorder []int, postStart, postEnd int) *TreeNode {
	if inStart > inEnd {
		return nil
	}
	root := &TreeNode{Val: postorder[postEnd]}
	var idx int
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == postorder[postEnd] {
			idx = i
			break
		}
	}
	leftSize := idx - inStart
	root.Left = build(inorder, inStart, idx-1, postorder, postStart, postStart+leftSize-1)
	root.Right = build(inorder, idx+1, inEnd, postorder, postStart+leftSize, postEnd-1)
	return root
}
