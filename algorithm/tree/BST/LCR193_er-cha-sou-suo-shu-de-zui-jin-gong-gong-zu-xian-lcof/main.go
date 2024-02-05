package main

/* 二叉搜索树的最近公共祖先 */

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getPath(root, target *TreeNode) []*TreeNode {
	path := []*TreeNode{}
	node := root
	for node != target {
		path = append(path, node)
		if node.Val < target.Val {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	path = append(path, node)
	return path
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pathP := getPath(root, p)
	pathQ := getPath(root, q)
	var ans *TreeNode
	for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i]; i++ {
		ans = pathP[i]
	}
	return ans
}
