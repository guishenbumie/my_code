package main

import "fmt"

/* 反转链表 */

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) ToString() string {
	if node == nil {
		return "nil"
	}
	s := ""
	curr := node
	for curr != nil {
		s = s + fmt.Sprint(curr.Val)
		curr = curr.Next
		if curr != nil {
			s = s + "->"
		}
	}
	return s
}

func initList(arr []int) *ListNode {
	if len(arr) <= 0 {
		return nil
	}
	head := &ListNode{}
	last := &ListNode{}
	for k, v := range arr {
		temp := &ListNode{
			Val:  v,
			Next: nil,
		}
		if k == 0 {
			last = temp
			head = temp
		} else {
			last.Next = temp
			last = temp
		}
	}
	return head
}

func main() {
	node11 := initList([]int{1, 2, 3})
	fmt.Println(node11.ToString())
	node12 := reverseList(node11)
	fmt.Println(node12.ToString())

	node21 := initList([]int{1, 2, 2, 1})
	fmt.Println(isPalindrome(node21))
	node22 := initList([]int{1, 2})
	fmt.Println(isPalindrome(node22))
	node23 := initList([]int{1, 2, 1})
	fmt.Println(isPalindrome(node23))
}

func reverseList(head *ListNode) *ListNode {
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

// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	arr := []int{}
	curr := head
	for curr != nil {
		arr = append(arr, curr.Val)
		curr = curr.Next
	}
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-i-1] {
			return false
		}
	}
	return true
}
