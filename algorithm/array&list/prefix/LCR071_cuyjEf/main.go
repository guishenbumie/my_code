package main

import (
	"math/rand"
	"sort"
)

/* 按权重随机选择 */

// 和528题一样
func main() {

}

type Solution struct {
	pre []int
}

func Constructor(w []int) Solution {
	pre := make([]int, len(w))
	pre[0] = w[0]
	for i := 1; i < len(w); i++ {
		pre[i] = pre[i-1] + w[i]
	}
	return Solution{pre}
}

func (this *Solution) PickIndex() int {
	x := rand.Intn(this.pre[len(this.pre)-1]) + 1
	return sort.SearchInts(this.pre, x)
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
