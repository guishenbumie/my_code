package main

import "fmt"

/* 反转链表 II */

//给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
//
//
//示例 1：
//
//
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
//示例 2：
//
//输入：head = [5], left = 1, right = 1
//输出：[5]
//
//
//提示：
//
//链表中节点数目为 n
//1 <= n <= 500
//-500 <= Node.val <= 500
//1 <= left <= right <= n

func main() {
	fmt.Println(reverseBetween(Constructor([]int{1, 2, 3, 4, 5}), 2, 4).ToString())
	fmt.Println(reverseBetween(Constructor([]int{5}), 1, 1).ToString())
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

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		l, _ := reverseN(head, right)
		return l
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

func reverseN(head *ListNode, right int) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}

	if right == 1 {
		return head, head.Next
	}

	last, suc := reverseN(head.Next, right-1)
	head.Next.Next = head
	head.Next = suc

	return last, suc
}
