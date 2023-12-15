package main

/* 反转二叉树的奇数层 */

//给你一棵 完美 二叉树的根节点 root ，请你反转这棵树中每个 奇数 层的节点值。
//
//例如，假设第 3 层的节点值是 [2, 1, 3,4, 7, 11, 29, 18] ，那么反转后它应该变成 [18, 29, 11, 7, 4, 3, 1, 2] 。
//反转后，返回树的根节点。
//
//完美 二叉树需满足：二叉树的所有父节点都有两个子节点，且所有叶子节点都在同一层。
//
//节点的 层数 等于该节点到根节点之间的边数。
//
//
//
//示例 1：
//
//2                        2
///     \                      /   \
//3       5 = >        5        3
///   \   /   \              /   \   /   \
//8     13    21   34            8      13 21    34
//
//输入：root = [2, 3, 5,8, 13, 21, 34]
//输出：[2, 5, 3, 8,13, 21, 34]
//解释：
//这棵树只有一个奇数层。
//在第 1 层的节点分别是 3、5 ，反转后为 5、3 。
//
//示例 2：
//
//7                    7
///      \ = >      /   \
//13      11            11       13
//
//输入：root = [7, 13, 11]
//输出：[7, 11, 13]
//解释：
//在第 1 层的节点分别是 13、11 ，反转后为 11、13 。
//示例 3：
//
//输入：root = [0, 1, 2,0, 0, 0, 0, 1,1, 1, 1, 2, 2,2, 2]
//输出：[0,2, 1, 0, 0, 0,0, 2, 2, 2, 2,1, 1, 1, 1]
//解释：奇数层由非零值组成。
//在第 1 层的节点分别是 1、2 ，反转后为 2、1 。
//在第 3 层的节点分别是 1、1、1、1、2、2、2、2 ，反转后为 2、2、2、2、1、1、1、1 。

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	q := []*TreeNode{root}
	isOdd := 0
	for len(q) > 0 {
		if isOdd != 0 {
			n := len(q)
			for i := 0; i < n/2; i++ {
				nodex, nodey := q[i], q[n-1-i]
				nodex.Val, nodey.Val = nodey.Val, nodex.Val
			}
		}
		tmp := make([]*TreeNode, 0, len(q)*2)
		for _, node := range q {
			if node.Left != nil {
				tmp = append(tmp, node.Left, node.Right)
			}
		}
		q = tmp
		isOdd ^= 1
	}
	return root
}
