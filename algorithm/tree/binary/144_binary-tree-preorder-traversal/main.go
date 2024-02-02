package main

/* 二叉树的前序遍历 */

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

// //递归方式
// func preorderTraversal(root *TreeNode) []int {
//     ans := []int{}
//     var traverse func(root *TreeNode)
//     traverse = func(root *TreeNode) {
//         if root == nil {
//             return
//         }
//         ans = append(ans, root.Val)
//         traverse(root.Left)
//         traverse(root.Right)
//     }
//     traverse(root)
//     return ans
// }

// 迭代方式
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	stk := []*TreeNode{}
	node := root
	for node != nil || len(stk) > 0 {
		for node != nil {
			ans = append(ans, node.Val)
			stk = append(stk, node)
			node = node.Left
		}
		node = stk[len(stk)-1].Right
		stk = stk[:len(stk)-1]
	}
	return ans
}
