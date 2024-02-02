package main

import (
	"fmt"
	"math"
)

/* 二叉树的最大深度 */

func main() {
	r1 := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(maxDepth(r1))

	r2 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(maxDepth(r2))

	r3 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: nil,
	}
	fmt.Println(maxDepth(r3))
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

func maxDepth(root *TreeNode) int {
	var traverse func(root *TreeNode) int
	traverse = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftMax := traverse(root.Left)
		rightMax := traverse(root.Right)
		//max := rightMax
		//if leftMax > rightMax {
		//	max = leftMax
		//}
		max := int(math.Max(float64(leftMax), float64(rightMax)))
		return max + 1
	}
	return traverse(root)
}

//func maxDepth(root *TreeNode) int {
//	depth, ans := 0, 0
//	var traverse func(root *TreeNode)
//	traverse = func(root *TreeNode) {
//		if root == nil {
//			return
//		}
//		depth++
//		ans = int(math.Max(float64(ans), float64(depth)))
//		traverse(root.Left)
//		traverse(root.Right)
//		depth--
//	}
//	return ans
//}
