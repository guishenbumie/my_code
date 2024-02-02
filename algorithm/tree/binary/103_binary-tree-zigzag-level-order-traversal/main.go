package main

/* 二叉树的锯齿形层序遍历 */

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

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stk := []*TreeNode{}
	stk = append(stk, root)
	ans := [][]int{}
	count := 0
	for len(stk) > 0 {
		sz := len(stk)
		arr := make([]int, sz)
		for i := 0; i < sz; i++ {
			node := stk[i]
			if count%2 == 0 {
				arr[i] = node.Val
			} else {
				arr[sz-i-1] = node.Val
			}
			if node.Left != nil {
				stk = append(stk, node.Left)
			}
			if node.Right != nil {
				stk = append(stk, node.Right)
			}
		}
		stk = stk[sz:]
		ans = append(ans, arr)
		count++
	}
	return ans
}
