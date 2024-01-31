package main

/* 图书整理 I */

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBookList(head *ListNode) []int {
	newHead := reverse(head)
	ans := []int{}
	for newHead != nil {
		ans = append(ans, newHead.Val)
		newHead = newHead.Next
	}
	return ans
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
