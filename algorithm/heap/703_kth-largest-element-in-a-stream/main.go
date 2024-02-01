package main

import "container/heap"

/* 数据流中的第 K 大元素 */

func main() {

}

type KthLargest struct {
	k  int
	mp *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	mp := NewMinHeap()
	heap.Init(mp)
	for i := 0; i < len(nums); i++ {
		heap.Push(mp, nums[i])
		if mp.Len() > k {
			heap.Pop(mp)
		}
	}
	return KthLargest{k, mp}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.mp, val)
	if this.mp.Len() > this.k {
		heap.Pop(this.mp)
	}
	return this.mp.Top()
}

type MinHeap struct {
	nums []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (h *MinHeap) Len() int {
	return len(h.nums)
}

func (h *MinHeap) Less(i, j int) bool {
	return h.nums[i] < h.nums[j]
}

func (h *MinHeap) Swap(i, j int) {
	h.nums[i], h.nums[j] = h.nums[j], h.nums[i]
}

func (h *MinHeap) Push(x interface{}) {
	h.nums = append(h.nums, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	x := h.nums[len(h.nums)-1]
	h.nums = h.nums[:len(h.nums)-1]
	return x
}

func (h *MinHeap) Top() int {
	return h.nums[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
