package main

import "fmt"

/* 删除链表的倒数第 N 个结点 */

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//
//
//示例 1：
//1->2->3->4->5
//1->2->3->5
//
//输入：head = [1, 2, 3, 4, 5], n = 2
//输出：[1,2, 3, 5]
//示例 2：
//
//输入：head = [1], n = 1
//输出：[]
//示例 3：
//
//输入：head = [1, 2], n = 1
//输出：[1]
//
//
//提示：
//
//链表中结点的数目为 sz
//1 <= sz <= 30
//0 <= Node.val <= 100
//1 <= n <= sz

func main() {
	l1 := Constructor([]int{1, 2, 3, 4, 5})
	fmt.Println(removeNthFromEnd(l1, 2).ToString())
	l2 := Constructor([]int{1})
	fmt.Println(removeNthFromEnd(l2, 1).ToString())
	l3 := Constructor([]int{1, 2})
	fmt.Println(removeNthFromEnd(l3, 1).ToString())
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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := &ListNode{
		Val:  -1,
		Next: head,
	}
	l := findFromEnd(temp, n+1)
	l.Next = l.Next.Next
	return temp.Next
}

// 返回链表的倒数第n个点
//
// 暴力：
// 遍历两次，第一次遍历整个list，得到list的总长度length；
// 倒数第n个点就是length-n，所以找到倒数第n个点的需要再遍历一次找到length-n就是目标元素；
//
// 更好的方式是双指针方法：
// 第一个指针遍历到n
// 第二个指针和第一个指针再同时遍历，直到第一个指针走到链表尾部null，第二个指针所在的位置就是length-n的目标元素所在位置
func findFromEnd(head *ListNode, n int) *ListNode {
	l1 := head
	for i := 0; i < n; i++ {
		l1 = l1.Next
	}
	l2 := head
	for l1 != nil {
		l2 = l2.Next
		l1 = l1.Next
	}
	return l2
}
