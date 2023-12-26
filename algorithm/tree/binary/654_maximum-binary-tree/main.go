package main

import "math"

/* 最大二叉树 */

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

func build(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	idx := 0
	max := math.MinInt32
	for i := l; i <= r; i++ {
		if nums[i] > max {
			max = nums[i]
			idx = i
		}
	}
	root := &TreeNode{Val: max}
	root.Left = build(nums, l, idx-1)
	root.Right = build(nums, idx+1, r)
	return root
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}
