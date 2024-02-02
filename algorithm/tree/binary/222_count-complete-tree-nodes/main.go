package main

/* 完全二叉树的节点个数 */

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
// func countNodes(root *TreeNode) int {
//     count := 0
//     var traverse func(root *TreeNode)
//     traverse = func(root *TreeNode) {
//         if root == nil {
//             return
//         }
//         traverse(root.Left)
//         traverse(root.Right)
//         count++
//     }
//     traverse(root)
//     return count
// }

// 迭代方式
func countNodes(root *TreeNode) int {
	count := 0
	if root == nil {
		return 0
	}
	stk := []*TreeNode{}
	stk = append(stk, root)
	for len(stk) > 0 {
		sz := len(stk)
		for i := 0; i < sz; i++ {
			node := stk[i]
			if node.Left != nil {
				stk = append(stk, node.Left)
			}
			if node.Right != nil {
				stk = append(stk, node.Right)
			}
		}
		stk = stk[sz:]
		count += sz
	}
	return count
}
