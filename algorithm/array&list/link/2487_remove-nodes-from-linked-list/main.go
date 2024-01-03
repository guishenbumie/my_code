package main

import "fmt"

/* 从链表中移除节点 */

func main() {
	l1 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 13,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Println(removeNodes(l1).ToString())
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

func (l *ListNode) ToString() string {
	var str string
	for curr := l; curr != nil; curr = curr.Next {
		str += fmt.Sprintf("%d ", curr.Val)
	}
	return str
}

func removeNodes(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeNodes(head.Next)
	if head.Next != nil && head.Next.Val > head.Val {
		return head.Next
	} else {
		return head
	}
}
