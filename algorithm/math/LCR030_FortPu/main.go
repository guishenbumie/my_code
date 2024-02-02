package main

import "math/rand"

/* 时间插入、删除和获取随机元素 */

// 和380题一样
func main() {

}

type RandomizedSet struct {
	nums []int
	m    map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: make([]int, 0),
		m:    make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.m[val] = len(this.nums) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if idx, ok := this.m[val]; !ok {
		return false
	} else {
		last := len(this.nums) - 1
		this.nums[idx] = this.nums[last]
		this.m[this.nums[idx]] = idx
		this.nums = this.nums[:last]
		delete(this.m, val)
		return true
	}
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
