package main

/* 有序链表转换二叉搜索树 */

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	var traverse func(l, r *ListNode) *TreeNode
	traverse = func(l, r *ListNode) *TreeNode {
		if l == r {
			return nil
		}
		mid := getMedian(l, r)
		root := &TreeNode{Val: mid.Val}
		root.Left = traverse(l, mid)
		root.Right = traverse(mid.Next, r)
		return root
	}
	return traverse(head, nil)
}

func getMedian(l, r *ListNode) *ListNode {
	slow, fast := l, l
	for fast != r && fast.Next != r {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
