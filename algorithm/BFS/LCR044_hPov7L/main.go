package main

import "math"

/*在每个树行中找最大值*/

// 和515题一样
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

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{}

	q := []*TreeNode{}
	q = append(q, root)
	for len(q) > 0 {
		sz := len(q)
		max := math.MinInt32
		for i := 0; i < sz; i++ {
			node := q[i]
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[sz:]
		ans = append(ans, max)
	}

	return ans
}
