package main

import "fmt"

/* 二叉搜索树中第K小的元素 */

func main() {
	r1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(kthSmallest(r1, 1))

	r2 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(kthSmallest(r2, 3))
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

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	var res int
	var rank int
	var f func(r *TreeNode, k int)
	f = func(r *TreeNode, k int) {
		if r == nil {
			return
		}
		f(r.Left, k)
		rank++
		if k == rank {
			res = r.Val
			return
		}
		f(r.Right, k)
	}
	f(root, k)
	return res
}
