package main

/* 路径总和 */

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
// func hasPathSum(root *TreeNode, targetSum int) bool {
//     var traverse func(root *TreeNode, targetSum int) bool
//     traverse = func(root *TreeNode, targetSum int) bool {
//         if root == nil {
//             return false
//         }
//         if root.Left == nil && root.Right == nil {
//             return root.Val == targetSum
//         }
//         diff := targetSum - root.Val
//         return hasPathSum(root.Left, diff) || hasPathSum(root.Right, diff)
//     }
//     return traverse(root, targetSum)
// }

// 广度优先
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	nodeStk := []*TreeNode{}
	valStk := []int{}
	nodeStk = append(nodeStk, root)
	valStk = append(valStk, root.Val)
	for len(nodeStk) > 0 {
		node := nodeStk[0]
		nodeStk = nodeStk[1:]
		val := valStk[0]
		valStk = valStk[1:]
		if node.Left == nil && node.Right == nil {
			if val == targetSum {
				return true
			}
			continue
		}
		if node.Left != nil {
			nodeStk = append(nodeStk, node.Left)
			valStk = append(valStk, node.Left.Val+val)
		}
		if node.Right != nil {
			nodeStk = append(nodeStk, node.Right)
			valStk = append(valStk, node.Right.Val+val)
		}
	}
	return false
}
