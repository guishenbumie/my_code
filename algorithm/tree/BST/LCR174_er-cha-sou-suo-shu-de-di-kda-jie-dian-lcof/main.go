package main

/* 寻找二叉搜索树中的目标节点 */

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

func findTargetNode(root *TreeNode, cnt int) int {
	ans := 0
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) { //中序遍历的变种，倒着中序，右中左
		if root == nil {
			return
		}
		traverse(root.Right)
		cnt--
		if cnt == 0 {
			ans = root.Val
		}
		traverse(root.Left)
	}
	traverse(root)
	return ans
}
