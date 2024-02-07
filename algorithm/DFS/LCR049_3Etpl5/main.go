package main

/* 求根节点到叶节点数字之和 */

// 和129题一样
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

func sumNumbers(root *TreeNode) int {
	ans := 0
	path := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		path = append(path, root.Val)
		defer func() {
			path = path[:len(path)-1]
		}()

		if root.Left == nil && root.Right == nil {
			num := 0
			for i := 0; i < len(path); i++ {
				num = num*10 + path[i]
			}
			ans += num
			return
		}

		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return ans
}
