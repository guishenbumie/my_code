package main

/* 彩灯装饰记录 II */

// 和102题一样
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
	for len(q) > 0 {
		sz := len(q)
		arr := make([]int, sz)
		for i := 0; i < sz; i++ {
			node := q[i]
			arr[i] = node.Val
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[sz:]
		ans = append(ans, append([]int(nil), arr...))
	}
	return ans
}
