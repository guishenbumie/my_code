package main

import "math/rand"

/* 链表随机节点 */

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

type Solution struct {
	arr []*ListNode
}

func Constructor(head *ListNode) Solution {
	arr := make([]*ListNode, 0)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	return Solution{arr: arr}
}

func (this *Solution) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))].Val
}

// 也可用水塘抽样算法
func getRandom(head *ListNode) int {
	var res int
	var cnt int
	for head != nil {
		cnt++
		if rand.Intn(cnt) == 0 {
			res = head.Val
		}
		head = head.Next
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
