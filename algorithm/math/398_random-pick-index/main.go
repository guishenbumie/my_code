package main

import "math/rand"

/* 随机数索引 */

func main() {

}

type Solution struct {
	m map[int][]int
}

func Constructor(nums []int) Solution {
	m := make(map[int][]int)
	for i, v := range nums {
		m[v] = append(m[v], i)
	}
	return Solution{m: m}
}

func (this *Solution) Pick(target int) int {
	v := this.m[target]
	return v[rand.Intn(len(v))]
}

// 也可用水塘抽样算法
func getRandom(nums []int, target int) int {
	var res int
	var cnt int
	for i, num := range nums {
		if num == target {
			cnt++
			if rand.Intn(cnt) == 0 {
				res = i
			}
		}
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
