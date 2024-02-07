package main

/* 二叉树中和为目标值的路径 */

// 和113题一样
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

func pathTarget(root *TreeNode, target int) [][]int {
	ans := [][]int{}
	path := []int{}
	sum := 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		sum += root.Val
		defer func() {
			path = path[:len(path)-1]
			sum -= root.Val
		}()
		if root.Left == nil && root.Right == nil {
			if sum == target {
				ans = append(ans, append([]int(nil), path...))
			}
		}
		dfs(root.Left)
		dfs(root.Right)

	}
	dfs(root)
	return ans
}
