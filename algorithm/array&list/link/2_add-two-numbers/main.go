package main

/* 两数相加 */

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	l := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		val := carry
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		carry = val / 10
		val = val % 10
		l.Next = &ListNode{Val: val}
		l = l.Next
	}
	return dummy.Next
}
