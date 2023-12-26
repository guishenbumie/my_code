package main

import "fmt"

/* 翻转二叉树 */

//给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
//
//
//
//示例 1：
//		  4						  4
//		/	\					/	\
//	2			7			7			2
//  /   \		  /	   \	  /   \       /   \
//1		3	6		9	9		6	3		1
//
//输入：root = [4,2,7,1,3,6,9]
//输出：[4,7,2,9,6,3,1]
//示例 2：
//	2			2
//  /   \	      /   \
//1		3	3		1
//
//输入：root = [2,1,3]
//输出：[2,3,1]
//示例 3：
//
//输入：root = []
//输出：[]
//
//
//提示：
//
//树中节点数目范围在 [0, 100] 内
//-100 <= Node.val <= 100

func main() {
	n := Constructor([]int{4, 2, 7, 1, 3, 6, 9}, 0)
	PrintTree(n)
	nn := invertTree(n)
	fmt.Println()
	PrintTree(nn)
	//fmt.Println(nn)
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

// 初始化一个二叉树
func Constructor(nums []int, i int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	root := &TreeNode{Val: nums[i]}
	if i < len(nums) && 2*i+1 < len(nums) {
		root.Left = Constructor(nums, 2*i+1)
	}
	if i < len(nums) && 2*i+2 < len(nums) {
		root.Right = Constructor(nums, 2*i+2)
	}
	return root
}

func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val)
	fmt.Print("  ")
	PrintTree(root.Left)
	PrintTree(root.Right)
}

//func (n *TreeNode) Print(root *TreeNode) {
//
//}
//
//func (n *TreeNode) ToArr(root *TreeNode) []int {
//	arr := make([]int, 0)
//	if root == nil {
//		return arr
//	}
//	arr = append(arr, root.Val)
//	arr = append(arr, n.ToArr(root.Left)...)
//	arr = append(arr, n.ToArr(root.Right)...)
//	return arr
//}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
