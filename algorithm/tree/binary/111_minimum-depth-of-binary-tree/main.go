package main

import (
	"container/list"
	"fmt"
)

/* 二叉树的最小深度 */

func main() {
	r1 := &TreeNode{
		Val: 3,
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
	fmt.Println(minDepth1(r1))

	r2 := &TreeNode{
		Val:  2,
		Left: nil,
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:  4,
				Left: nil,
				Right: &TreeNode{
					Val:  5,
					Left: nil,
					Right: &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}
	fmt.Println(minDepth1(r2))
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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := list.New()
	q.PushBack(root)
	step := 1
	for q.Len() > 0 {
		sz := q.Len()
		for i := 0; i < sz; i++ {
			curr := q.Remove(q.Front()).(*TreeNode)
			//判断是否到达终点
			if curr.Left == nil && curr.Right == nil {
				return step
			}
			//将相邻节点加入队列
			if curr.Left != nil {
				q.PushBack(curr.Left)
			}
			if curr.Right != nil {
				q.PushBack(curr.Right)
			}
		}
		step++
	}
	return step
}

func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := make([]*TreeNode, 1)
	q[0] = root
	step := 1
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			curr := q[0]
			q = q[1:]
			//判断是否到达终点
			if curr.Left == nil && curr.Right == nil {
				return step
			}
			//将相邻节点加入队列
			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
		step++
	}
	return step
}
