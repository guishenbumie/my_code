package main

import "fmt"

/* 删除排序链表中的重复元素 */

//给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
//
//
//
//示例 1：
//1->1->2  =>  1->2
//
//输入：head = [1,1,2]
//输出：[1,2]
//
//示例 2：
//1->1->2->3->3  =>  1->2->3
//
//输入：head = [1,1,2,3,3]
//输出：[1,2,3]
//
//
//提示：
//
//链表中节点数目在范围 [0, 300] 内
//-100 <= Node.val <= 100
//题目数据保证链表已经按升序 排列

func main() {
	//l := Constructor([]int{1, 2, 3})
	//fmt.Println(l.ToString())
	fmt.Println(deleteDuplicates(Constructor([]int{1, 1, 2})).ToString())
	fmt.Println(deleteDuplicates(Constructor([]int{1, 1, 2, 3, 3})).ToString())
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

	head := &ListNode{}
	last := &ListNode{}
	for i := 0; i < len(nums); i++ {
		temp := &ListNode{
			Val:  nums[i],
			Next: nil,
		}
		if i == 0 {
			head = temp
			last = temp
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

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}
