package main

import "fmt"

/* 链表的中间结点 */

//给你单链表的头结点 head ，请你找出并返回链表的中间结点。
//
//如果有两个中间结点，则返回第二个中间结点。
//
//
//
//示例 1：
//
//
//输入：head = [1,2,3,4,5]
//输出：[3,4,5]
//解释：链表只有一个中间结点，值为 3 。
//示例 2：
//
//
//输入：head = [1,2,3,4,5,6]
//输出：[4,5,6]
//解释：该链表有两个中间结点，值分别为 3 和 4 ，返回第二个结点。
//
//
//提示：
//
//链表的结点数范围是 [1, 100]
//1 <= Node.val <= 100

func main() {
	fmt.Println(middleNode(Constructor([]int{1, 2, 3, 4, 5})).ToString())
	fmt.Println(middleNode(Constructor([]int{1, 2, 3, 4, 5, 6})).ToString())
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

// 使用快慢指针方式
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

/* 延伸1 - 判断一个链表是否包含环 */
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		//快慢指针最终相遇，有环
		if slow == fast {
			return true
		}
	}
	return false
}

/* 延伸2 - 一个链表含环，返回环的起点 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	//说明没有环
	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
