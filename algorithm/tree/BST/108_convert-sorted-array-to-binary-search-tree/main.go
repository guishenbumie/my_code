package main

/* 将有序数组转换为二叉搜索树 */

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

func sortedArrayToBST(nums []int) *TreeNode {
	var traverse func(nums []int, left, right int) *TreeNode
	traverse = func(nums []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + (right-left)/2
		root := &TreeNode{Val: nums[mid]}
		root.Left = traverse(nums, left, mid-1)
		root.Right = traverse(nums, mid+1, right)
		return root
	}
	return traverse(nums, 0, len(nums)-1)
}
