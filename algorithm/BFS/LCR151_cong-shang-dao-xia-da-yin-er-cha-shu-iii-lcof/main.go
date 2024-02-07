package main

/* 彩灯装饰记录 III */

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

func decorateRecord(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	count := 0
	for len(q) > 0 {
		sz := len(q)
		arr := make([]int, sz)
		for i := 0; i < sz; i++ {
			node := q[i]
			if count%2 == 0 {
				arr[i] = node.Val
			} else {
				arr[sz-i-1] = node.Val
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[sz:]
		ans = append(ans, arr)
		count++
	}
	return ans
}
