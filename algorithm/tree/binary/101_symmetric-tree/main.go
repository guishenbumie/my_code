package main

/* 对称二叉树 */

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
// func isSymmetric(root *TreeNode) bool {
//     var traverse func(n1, n2 *TreeNode) bool
//     traverse = func(n1, n2 *TreeNode) bool {
//         if n1 == nil && n2 == nil {
//             return true
//         }
//         if n1 == nil || n2 == nil {
//             return false
//         }
//         if n1.Val != n2.Val {
//             return false
//         }
//         return traverse(n1.Left, n2.Right) && traverse(n1.Right, n2.Left)
//     }
//     return traverse(root, root)
// }

// 迭代方式
func isSymmetric(root *TreeNode) bool {
	n1, n2 := root, root
	q := []*TreeNode{}
	q = append(q, n1)
	q = append(q, n2)
	for len(q) > 0 {
		n1, n2 = q[0], q[1]
		q = q[2:]
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}

		q = append(q, n1.Left)
		q = append(q, n2.Right)

		q = append(q, n1.Right)
		q = append(q, n2.Left)
	}
	return true
}
