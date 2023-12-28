package main

import (
	"fmt"
	"math"
)

/* 二叉搜索子树的最大键值和 */

func main() {
	//root := &TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val:   4,
	//		Left:  &TreeNode{Val: 2},
	//		Right: &TreeNode{Val: 4},
	//	},
	//	Right: &TreeNode{
	//		Val:  3,
	//		Left: &TreeNode{Val: 2},
	//		Right: &TreeNode{
	//			Val:   5,
	//			Left:  &TreeNode{Val: 4},
	//			Right: &TreeNode{Val: 6},
	//		},
	//	},
	//}
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		},
	}
	fmt.Println(maxSumBST(root))
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

var maxSum int

func maxSumBST(root *TreeNode) int {
	maxSum = 0
	traverse(root)
	return maxSum
}

// 返回值，是否是BST树、root为根所有节点中最小值、root为根所有节点中最大值、root为根所有节点的和
func traverse(root *TreeNode) (bool, int, int, int) {
	if root == nil {
		return true, math.MaxInt32, math.MinInt32, 0
	}
	leftIsBST, leftMin, leftMax, leftSum := traverse(root.Left)
	rightIsBST, rightMin, rightMax, rightSum := traverse(root.Right)

	var isBST bool
	var min int
	var max int
	var sum int
	if leftIsBST && rightIsBST && root.Val > leftMax && root.Val < rightMin {
		isBST = true
		if leftMin > root.Val {
			min = root.Val
		} else {
			min = leftMin
		}
		if rightMax > root.Val {
			max = rightMax
		} else {
			max = root.Val
		}
		sum = leftSum + rightSum + root.Val
		if maxSum < sum {
			maxSum = sum
		}
	} else {
		isBST = false
	}
	return isBST, min, max, sum
}
