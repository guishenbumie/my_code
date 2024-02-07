package main

import (
	"fmt"
	"math"
)

/* 二叉树中的最大路径和 */

func main() {
	r1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(maxPathSum(r1))

	r2 := &TreeNode{
		Val: -10,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(maxPathSum(r2))

	r3 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   -1,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	fmt.Println(maxPathSum(r3))
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

func maxPathSum(root *TreeNode) int {
	res := math.MinInt32
	var f func(r *TreeNode) int
	f = func(r *TreeNode) int {
		if r == nil {
			return 0
		}
		leftMax := max(f(r.Left), 0)
		rightMax := max(f(r.Right), 0)
		temp := leftMax + rightMax + r.Val
		res = max(temp, res)
		return r.Val + max(leftMax, rightMax)
	}
	f(root)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
