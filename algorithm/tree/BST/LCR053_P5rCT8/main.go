package main

/* 二叉搜索树中的中序后继 */

// 和285题一样
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

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	stk := []*TreeNode{}
	var prev *TreeNode
	curr := root
	for len(stk) > 0 || curr != nil {
		for curr != nil {
			stk = append(stk, curr)
			curr = curr.Left
		}
		curr = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if prev == p {
			return curr
		}
		prev = curr
		curr = curr.Right
	}
	return nil
}
