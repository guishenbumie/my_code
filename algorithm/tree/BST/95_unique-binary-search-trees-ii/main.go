package main

/* 不同的二叉搜索树 II */

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

func generateTrees(n int) []*TreeNode {
	var build func(lo, hi int) []*TreeNode
	build = func(lo, hi int) []*TreeNode {
		ans := []*TreeNode{}
		if lo > hi {
			ans = append(ans, nil)
			return ans
		}
		for i := lo; i <= hi; i++ {
			leftTree := build(lo, i-1)
			rightTree := build(i+1, hi)
			for _, left := range leftTree {
				for _, right := range rightTree {
					root := &TreeNode{
						Val:   i,
						Left:  left,
						Right: right,
					}
					ans = append(ans, root)
				}
			}
		}
		return ans
	}
	return build(1, n)
}
