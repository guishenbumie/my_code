package main

import "fmt"

/* 填充每个节点的下一个右侧节点指针 */

//给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。二叉树定义如下：
//
//struct Node {
//int val;
//Node *left;
//Node *right;
//Node *next;
//}
//填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
//
//初始状态下，所有 next 指针都被设置为 NULL。
//
//
//
//示例 1：
//		  1						  1	-> nil
//	   /     \			       /	  \
//	2			3			2	  ->	3 -> nil
//  /   \       /    \	  /   \       /   \
//4		5	6		7	4	->	5 ->6	->	7 -> nil
//
//输入：root = [1,2,3,4,5,6,7]
//输出：[1,#,2,3,#,4,5,6,7,#]
//解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。序列化的输出按层序遍历排列，同一层节点由 next 指针连接，'#' 标志着每一层的结束。
//示例 2:
//
//输入：root = []
//输出：[]
//
//
//提示：
//
//树中节点的数量在 [0, 212 - 1] 范围内
//-1000 <= node.val <= 1000
//
//
//进阶：
//
//你只能使用常量级额外空间。
//使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

func main() {
	n := Constructor([]int{1, 2, 3, 4, 5, 6, 7}, 0)
	PrintTree(n)
	nn := connect(n)
	fmt.Println()
	PrintTree(nn)
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Constructor(nums []int, i int) *Node {
	length := len(nums)
	if length <= 0 {
		return nil
	}
	root := &Node{Val: nums[i]}
	if i < length && 2*i+1 < length {
		root.Left = Constructor(nums, 2*i+1)
	}
	if 2 < length && 2*i+2 < length {
		root.Right = Constructor(nums, 2*i+2)
	}
	return root
}

func PrintTree(root *Node) {
	if root == nil {
		return
	}
	fmt.Print(root.Val)
	fmt.Print("  ")
	PrintTree(root.Left)
	PrintTree(root.Right)
}

// 前序遍历
func connectTwoNode(n1, n2 *Node) {
	if n1 == nil || n2 == nil {
		return
	}
	n1.Next = n2                      //将传入的两个连接起来
	connectTwoNode(n1.Left, n1.Right) //将n1的两个子节点连接起来
	connectTwoNode(n2.Left, n2.Right) //将n2的两个子节点连接起来
	connectTwoNode(n1.Right, n2.Left) //将非同父但同一层的挨着的两个子节点连接起来
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	connectTwoNode(root.Left, root.Right)
	return root
}
