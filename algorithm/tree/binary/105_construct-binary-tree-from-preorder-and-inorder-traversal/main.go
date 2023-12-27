package main

/* 从前序与中序遍历序列构造二叉树 */

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

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func build(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	root := &TreeNode{Val: preorder[preStart]}
	var idx int
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == preorder[preStart] {
			idx = i
			break
		}
	}
	leftSize := idx - inStart
	root.Left = build(preorder, preStart+1, preStart+leftSize, inorder, inStart, idx-1)
	root.Right = build(preorder, preStart+leftSize+1, preEnd, inorder, idx+1, inEnd)
	return root
}
