package main

/* 分隔链表 */

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

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left, right := &ListNode{}, &ListNode{}
	dummyLeft := left
	dummyRight := right
	curr := head
	for curr != nil {
		if curr.Val < x {
			left.Next = curr
			left = left.Next
		} else {
			right.Next = curr
			right = right.Next
		}
		curr = curr.Next
	}
	right.Next = nil
	left.Next = dummyRight.Next
	return dummyLeft.Next
}
