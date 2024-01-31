package main

import "container/heap"

/* 数据流中的中位数 */

// 和295题一样
func main() {

}

type MinHeap struct {
	nums []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (minH *MinHeap) Len() int {
	return len(minH.nums)
}

func (minH *MinHeap) Less(i, j int) bool {
	return minH.nums[i] < minH.nums[j]
}

func (minH *MinHeap) Swap(i, j int) {
	minH.nums[i], minH.nums[j] = minH.nums[j], minH.nums[i]
}

func (minH *MinHeap) Push(x interface{}) {
	minH.nums = append(minH.nums, x.(int))
}

func (minH *MinHeap) Pop() interface{} {
	x := minH.nums[len(minH.nums)-1]
	minH.nums = minH.nums[:len(minH.nums)-1]
	return x
}

func (minH *MinHeap) Top() int {
	return minH.nums[0]
}

type MaxHeap struct {
	nums []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (maxH *MaxHeap) Len() int {
	return len(maxH.nums)
}

func (maxH *MaxHeap) Less(i, j int) bool {
	return maxH.nums[i] > maxH.nums[j]
}

func (maxH *MaxHeap) Swap(i, j int) {
	maxH.nums[i], maxH.nums[j] = maxH.nums[j], maxH.nums[i]
}

func (maxH *MaxHeap) Push(x interface{}) {
	maxH.nums = append(maxH.nums, x.(int))
}

func (maxH *MaxHeap) Pop() interface{} {
	x := maxH.nums[len(maxH.nums)-1]
	maxH.nums = maxH.nums[:len(maxH.nums)-1]
	return x
}

func (maxH *MaxHeap) Top() int {
	return maxH.nums[0]
}

type MedianFinder struct {
	minH *MinHeap
	maxH *MaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	mf := MedianFinder{
		minH: NewMinHeap(),
		maxH: NewMaxHeap(),
	}
	heap.Init(mf.minH)
	heap.Init(mf.maxH)
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.maxH, num)
	heap.Push(this.minH, heap.Pop(this.maxH))
	for this.maxH.Len() < this.minH.Len() {
		heap.Push(this.maxH, heap.Pop(this.minH))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxH.Len() > this.minH.Len() {
		return float64(this.maxH.Top())
	}
	return float64(this.maxH.Top()+this.minH.Top()) * 0.5
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
