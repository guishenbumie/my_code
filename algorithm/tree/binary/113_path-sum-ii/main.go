package main

/* 路径总和 II */

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

// //深度优先
// func pathSum(root *TreeNode, targetSum int) [][]int {
//     ans := [][]int{}
//     path := []int{}
//     var traverse func(root *TreeNode, targetSum int)
//     traverse = func(root *TreeNode, targetSum int) {
//         if root == nil {
//             return
//         }
//         path = append(path, root.Val)
//         // defer func() {
//         //     path = path[:len(path)-1]
//         // }()
//         if root.Left == nil && root.Right == nil && targetSum == root.Val {
//             ans = append(ans, append([]int(nil), path...))
//             path = path[:len(path)-1]
//             return
//         }
//         diff := targetSum - root.Val
//         traverse(root.Left, diff)
//         traverse(root.Right, diff)
//         path = path[:len(path)-1]
//     }
//     traverse(root, targetSum)
//     return ans
// }

// 广度优先
type pair struct {
	node *TreeNode
	left int
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}

	parent := map[*TreeNode]*TreeNode{}

	getPath := func(node *TreeNode) []int {
		path := []int{}
		for node != nil {
			path = append(path, node.Val)
			node = parent[node]
		}
		for i, j := 0, len(path)-1; i < j; i++ {
			path[i], path[j] = path[j], path[i]
			j--
		}
		return path
	}

	ans := [][]int{}
	q := []pair{{root, targetSum}}
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		node := p.node
		left := p.left - node.Val
		if node.Left == nil && node.Right == nil {
			if left == 0 {
				ans = append(ans, getPath(node))
			}
		} else {
			if node.Left != nil {
				parent[node.Left] = node
				q = append(q, pair{node.Left, left})
			}
			if node.Right != nil {
				parent[node.Right] = node
				q = append(q, pair{node.Right, left})
			}
		}
	}
	return ans
}
