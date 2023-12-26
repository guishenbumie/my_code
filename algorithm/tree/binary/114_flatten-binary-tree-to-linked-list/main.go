package main

/* 二叉树展开为链表 */

//给你二叉树的根结点 root ，请你将它展开为一个单链表：
//
//展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
//展开后的单链表应该与二叉树 先序遍历 顺序相同。
//
//
//示例 1：
//
//
//输入：root = [1,2,5,3,4,null,6]
//输出：[1,null,2,null,3,null,4,null,5,null,6]
//示例 2：
//
//输入：root = []
//输出：[]
//示例 3：
//
//输入：root = [0]
//输出：[0]
//
//
//提示：
//
//树中结点数在范围 [0, 2000] 内
//-100 <= Node.val <= 100

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

func construct(nums []int, i int) *TreeNode {
	length := len(nums)
	if length <= 0 {
		return nil
	}
	root := &TreeNode{Val: nums[i]}
	if i < length && 2*i+1 < length {
		root.Left = construct(nums, 2*i+1)
	}
	if i < length && 2*i+2 < length {
		root.Right = construct(nums, 2*i+2)
	}
	return root
}

func Constructor(nums []int) *TreeNode {
	return construct(nums, 0)
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	l := root.Left
	r := root.Right
	root.Right = l
	root.Left = nil
	curr := root
	for curr.Right != nil {
		curr = curr.Right
	}
	curr.Right = r
}
