package main

/* 恢复二叉搜索树 */

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

func recoverTree(root *TreeNode) {
	var x, y, prev *TreeNode
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)

		if prev != nil && prev.Val > root.Val {
			y = root
			if x == nil {
				x = prev
			}
		}
		prev = root

		traverse(root.Right)
	}
	traverse(root)
	x.Val, y.Val = y.Val, x.Val
}

// func recoverTree(root *TreeNode)  {
//     stk := []*TreeNode{}
//     var x, y, prev *TreeNode
//     for len(stk) > 0 || root != nil {
//         for root != nil {
//             stk = append(stk, root)
//             root = root.Left
//         }
//         root = stk[len(stk)-1]
//         stk = stk[:len(stk)-1]
//         if prev != nil && root.Val < prev.Val {
//             y = root
//             if x == nil {
//                 x = prev
//             } else {
//                 break
//             }
//         }
//         prev = root
//         root = root.Right
//     }
//     x.Val, y.Val = y.Val, x.Val
// }
