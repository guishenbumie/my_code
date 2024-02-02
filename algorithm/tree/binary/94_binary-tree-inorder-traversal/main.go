package main

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

// //递归的方式
// func inorderTraversal(root *TreeNode) []int {
//     ans := []int{}
//     var traverse func(r *TreeNode)
//     traverse = func(r *TreeNode) {
//         if r == nil {
//             return
//         }
//         traverse(r.Left)
//         ans = append(ans, r.Val)
//         traverse(r.Right)
//     }
//     traverse(root)
//     return ans
// }

//func inorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//	ans := []int{}
//	stk := list.New()
//	for root != nil {
//		stk.PushBack(root)
//		root = root.Left
//	}
//	for stk.Len() > 0 {
//		curr := stk.Remove(stk.Back()).(*TreeNode)
//		ans = append(ans, curr.Val)
//		curr = curr.Right
//		for curr != nil {
//			stk.PushBack(curr)
//			curr = curr.Left
//		}
//	}
//	return ans
//}

// 迭代的方式
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ans := []int{}
	stk := []*TreeNode{}
	for root != nil {
		stk = append(stk, root)
		root = root.Left
	}
	for len(stk) > 0 {
		curr := stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		ans = append(ans, curr.Val)
		curr = curr.Right
		for curr != nil {
			stk = append(stk, curr)
			curr = curr.Left
		}
	}
	return ans
}
