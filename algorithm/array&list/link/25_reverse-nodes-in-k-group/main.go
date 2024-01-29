package main

/* K 个一组翻转链表 */

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

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	l1, l2 := head, head
	for i := 0; i < k; i++ {
		if l2 == nil {
			return head
		}
		l2 = l2.Next
	}
	newHead := reverse(l1, l2)
	l1.Next = reverseKGroup(l2, k)
	return newHead
}

// 反转[l1, l2)区间的元素，左闭右开
func reverse(l1, l2 *ListNode) *ListNode {
	var prev *ListNode
	curr, next := l1, l1
	for curr != l2 {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
