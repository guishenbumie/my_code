package main

/* 二叉树的后序遍历 */

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
// func postorderTraversal(root *TreeNode) []int {
//     ans := []int{}
//     var traverse func(root *TreeNode)
//     traverse = func(root *TreeNode) {
//         if root == nil {
//             return
//         }
//         traverse(root.Left)
//         traverse(root.Right)
//         ans = append(ans, root.Val)
//     }
//     traverse(root)
//     return ans
// }

// 迭代方式
func postorderTraversal(root *TreeNode) []int {
	ans := []int{}
	stk := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			ans = append(ans, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}
	}
	return ans
}
