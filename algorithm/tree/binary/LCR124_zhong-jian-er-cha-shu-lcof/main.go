package main

/* 推理二叉树 */

// 和105题一样
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

func deduceTree(preorder []int, inorder []int) *TreeNode {
	var build func(po []int, poStart, poEnd int, io []int, ioStart, ioEnd int) *TreeNode
	build = func(po []int, poStart, poEnd int, io []int, ioStart, ioEnd int) *TreeNode {
		if poStart > poEnd {
			return nil
		}
		root := &TreeNode{Val: preorder[poStart]}
		var idx int
		for i := ioStart; i <= ioEnd; i++ {
			if io[i] == root.Val {
				idx = i
				break
			}
		}
		length := idx - ioStart
		root.Left = build(po, poStart+1, poStart+length, io, ioStart, idx-1)
		root.Right = build(po, poStart+length+1, poEnd, io, idx+1, ioEnd)
		return root
	}
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}
