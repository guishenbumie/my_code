package main

/* 找树左下角的值 */

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

// 深度优先
func findBottomLeftValue(root *TreeNode) int {
	currHeight := 0
	ans := 0
	var dfs func(root *TreeNode, height int)
	dfs = func(root *TreeNode, height int) {
		if root == nil {
			return
		}
		height++
		dfs(root.Left, height)
		dfs(root.Right, height)
		if height > currHeight {
			currHeight = height
			ans = root.Val
		}
	}
	dfs(root, 0)
	return ans
}

//// 广度优先
//func findBottomLeftValue(root *TreeNode) int {
//	ans := 0
//
//	q := []*TreeNode{root}
//	for len(q) > 0 {
//		sz := len(q)
//		for i := 0; i < sz; i++ {
//			node := q[i]
//			ans = node.Val
//			if node.Right != nil {
//				q = append(q, node.Right)
//			}
//			if node.Left != nil {
//				q = append(q, node.Left)
//			}
//		}
//		q = q[sz:]
//	}
//
//	return ans
//}
