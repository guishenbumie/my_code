package main

/* 二叉树的层序遍历 II */

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

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	stk := []*TreeNode{}
	stk = append(stk, root)
	for len(stk) > 0 {
		sz := len(stk)
		arr := []int{}
		for i := 0; i < sz; i++ {
			node := stk[i]
			arr = append(arr, node.Val)
			if node.Left != nil {
				stk = append(stk, node.Left)
			}
			if node.Right != nil {
				stk = append(stk, node.Right)
			}
		}
		stk = stk[sz:]
		ans = append([][]int{arr}, ans...)
	}
	return ans
}
