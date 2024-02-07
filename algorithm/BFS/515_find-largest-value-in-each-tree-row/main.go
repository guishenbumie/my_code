package main

import (
	"container/list"
	"fmt"
	"math"
)

/* 在每个树行中找最大值 */

func main() {
	r1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(largestValues(r1))

	r2 := &TreeNode{
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
	fmt.Println(largestValues(r2))
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

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	l := list.New()
	l.PushBack(root)
	for l.Len() > 0 {
		max := math.MinInt32
		sz := l.Len()
		for i := 0; i < sz; i++ {
			curr := l.Remove(l.Front()).(*TreeNode)
			if curr.Val > max {
				max = curr.Val
			}
			if curr.Left != nil {
				l.PushBack(curr.Left)
			}
			if curr.Right != nil {
				l.PushBack(curr.Right)
			}
		}
		res = append(res, max)
	}
	return res
}
