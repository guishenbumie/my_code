package main

import "fmt"

/* 反转链表 */

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//
//
//示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]
//示例 2：
//
//
//输入：head = [1,2]
//输出：[2,1]
//示例 3：
//
//输入：head = []
//输出：[]
//
//
//提示：
//
//链表中节点的数目范围是 [0, 5000]
//-5000 <= Node.val <= 5000

func main() {
	fmt.Println(reverseList(Constructor([]int{1, 2, 3, 4, 5})).ToString())
	fmt.Println(reverseList(Constructor([]int{1, 2})).ToString())
	fmt.Println(reverseList(Constructor([]int{})).ToString())

	l, _ := reverseN(Constructor([]int{1, 2, 3, 4, 5, 6}), 3)
	fmt.Println(l.ToString())
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Constructor(nums []int) *ListNode {
	if len(nums) <= 0 {
		return nil
	}
	head, last := new(ListNode), new(ListNode)
	for i := 0; i < len(nums); i++ {
		temp := &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		if i == 0 {
			last = temp
			head = temp
		} else {
			last.Next = temp
			last = temp
		}
	}
	return head
}

func (h *ListNode) ToString() string {
	str := ""
	curr := h
	for curr != nil {
		str += fmt.Sprint(curr.Val)
		if curr.Next != nil {
			str += "->"
		}
		curr = curr.Next
	}
	return str
}

////迭代方式
//func reverseList(head *ListNode) *ListNode {
//	var prev *ListNode
//	curr := head
//	for curr != nil {
//		next := curr.Next
//		curr.Next = prev
//		prev = curr
//		curr = next
//	}
//	return prev
//}

// 递归方式
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// 反转前n个节点，n<=链表长度
// 比如：
// 1->2->3->4->5->6，n=3
// 反转后3->2->1->4->5->6
func reverseN(head *ListNode, n int) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}

	//反转一个元素，就是它本身
	if n == 1 {
		return head, head.Next //返回头节点和后驱节点
	}

	last, successor := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return last, successor
}

//func reverseN(head *ListNode, n int) (*ListNode, *ListNode) {
//	if head == nil {
//		return nil, nil
//	}
//	if n == 1 {
//		return head, head.Next
//	}
//	last, next := reverseN(head.Next, n-1)
//	head.Next.Next = head
//	head.Next = next
//	return last, next
//}
