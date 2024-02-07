package main

/* 二叉树的层序遍历 */

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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	stk := []*TreeNode{}
	stk = append(stk, root)
	for len(stk) > 0 {
		arr := []int{}
		sz := len(stk)
		for i := 0; i < sz; i++ {
			node := stk[0]
			stk = stk[1:]
			if node.Left != nil {
				stk = append(stk, node.Left)
			}
			if node.Right != nil {
				stk = append(stk, node.Right)
			}
			arr = append(arr, node.Val)
		}
		ans = append(ans, arr)
	}
	return ans
}
