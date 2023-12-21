package main

import (
	"container/heap"
	"fmt"
)

/* 合并 K 个升序链表 */

//给你一个链表数组，每个链表都已经按升序排列。
//
//请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
//示例 1：
//
//输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//1->4->5,
//1->3->4,
//2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//示例 2：
//
//输入：lists = []
//输出：[]
//示例 3：
//
//输入：lists = [[]]
//输出：[]
//
//
//提示：
//
//k == lists.length
//0 <= k <= 10^4
//0 <= lists[i].length <= 500
//-10^4 <= lists[i][j] <= 10^4
//lists[i] 按 升序 排列
//lists[i].length 的总和不超过 10^4

func main() {
	lists1 := []*ListNode{
		Constructor([]int{1, 4, 5}),
		Constructor([]int{1, 3, 4}),
		Constructor([]int{2, 6}),
	}
	fmt.Println(mergeKLists(lists1).ToString())
	fmt.Println(mergeKLists(nil).ToString())
	lists3 := []*ListNode{}
	fmt.Println(mergeKLists(lists3).ToString())
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

// 顺序合并
// 但最好的方式是使用优先级队列
//func mergeKLists(lists []*ListNode) *ListNode {
//	if lists == nil || len(lists) <= 0 {
//		return nil
//	}
//	if len(lists) == 1 {
//		return lists[0]
//	}
//	res := lists[0]
//	for i := 1; i < len(lists); i++ {
//		res = mergeList(res, lists[i])
//	}
//	return res
//}
//
//func mergeList(list1, list2 *ListNode) *ListNode {
//	head := &ListNode{
//		Val:  -1,
//		Next: nil,
//	}
//	p, p1, p2 := head, list1, list2
//	for p1 != nil && p2 != nil {
//		if p1.Val > p2.Val {
//			p.Next = p2
//			p2 = p2.Next
//		} else {
//			p.Next = p1
//			p1 = p1.Next
//		}
//		p = p.Next
//	}
//	if p1 != nil {
//		p.Next = p1
//	}
//	if p2 != nil {
//		p.Next = p2
//	}
//	return head.Next
//}

func mergeKLists(lists []*ListNode) *ListNode {
	h := hp{}
	for _, head := range lists {
		if head != nil {
			h = append(h, head)
		}
	}
	heap.Init(&h) // 堆化

	dummy := &ListNode{} // 哨兵节点，作为合并后链表头节点的前一个节点
	cur := dummy
	for len(h) > 0 { // 循环直到堆为空
		node := heap.Pop(&h).(*ListNode) // 剩余节点中的最小节点
		if node.Next != nil {            // 下一个节点不为空
			heap.Push(&h, node.Next) // 下一个节点有可能是最小节点，入堆
		}
		cur.Next = node // 合并到新链表中
		cur = cur.Next  // 准备合并下一个节点
	}
	return dummy.Next // 哨兵节点的下一个节点就是新链表的头节点
}

type hp []*ListNode

func (h hp) Len() int {
	return len(h)
}
func (h hp) Less(i, j int) bool {
	return h[i].Val < h[j].Val
} // 最小堆
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *hp) Push(v any) {
	*h = append(*h, v.(*ListNode))
}
func (h *hp) Pop() any {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}
