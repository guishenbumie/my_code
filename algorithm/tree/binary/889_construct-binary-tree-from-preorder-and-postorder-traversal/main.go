package main

/* 根据前序和后序遍历构造二叉树 */
//答案不唯一，前序+中序或者后序+中序可以唯一确定二叉树
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

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1)
}

func build(preorder []int, preStart, preEnd int, postorder []int, postStart, postEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	root := &TreeNode{Val: preorder[preStart]}
	if preStart == preEnd {
		return root
	}
	rootLeftVal := preorder[preStart+1]
	var idx int
	for i := postStart; i <= postEnd; i++ {
		if postorder[i] == rootLeftVal {
			idx = i
			break
		}
	}
	leftSize := idx - postStart + 1
	root.Left = build(preorder, preStart+1, preStart+leftSize, postorder, postStart, idx)
	root.Right = build(preorder, preStart+leftSize+1, preEnd, postorder, idx+1, postEnd-1)
	return root
}
