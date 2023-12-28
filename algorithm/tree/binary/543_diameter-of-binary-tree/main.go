package main

import "fmt"

/* 二叉树的直径 */

func main() {
	r1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(diameterOfBinaryTree(r1))

	r2 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	fmt.Println(diameterOfBinaryTree(r2))
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

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	var f func(root *TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftMax := f(root.Left)
		rightMax := f(root.Right)
		temp := leftMax + rightMax
		if temp > max {
			max = temp
		}
		if leftMax > rightMax {
			return leftMax + 1
		} else {
			return rightMax + 1
		}
	}
	f(root)
	return max
}
