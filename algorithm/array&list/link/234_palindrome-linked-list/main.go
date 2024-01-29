package main

/* 回文链表 */

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

////时间复杂度为O(n)，空间复杂度为O(n)
//func isPalindrome(head *ListNode) bool {
//	if head == nil {
//		return false
//	}
//	arr := []int{}
//	curr := head
//	for curr != nil {
//		arr = append(arr, curr.Val)
//		curr = curr.Next
//	}
//	for i := 0; i < len(arr)/2; i++ {
//		if arr[i] != arr[len(arr)-i-1] {
//			return false
//		}
//	}
//	return true
//}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	l1 := head
	l2 := reverse(slow)
	for l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
