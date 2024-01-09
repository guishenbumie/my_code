package main

import (
	"fmt"
	"math/rand"
)

/* 黑名单中的随机数 */

//给定一个整数 n 和一个 无重复 黑名单整数数组 blacklist 。设计一种算法，从 [0, n - 1] 范围内的任意整数中选取一个 未加入 黑名单 blacklist 的整数。任何在上述范围内且不在黑名单 blacklist 中的整数都应该有 同等的可能性 被返回。
//
//优化你的算法，使它最小化调用语言 内置 随机函数的次数。
//
//实现 Solution 类:
//
//Solution(int n, int[] blacklist) 初始化整数 n 和被加入黑名单 blacklist 的整数
//int pick() 返回一个范围为 [0, n - 1] 且不在黑名单 blacklist 中的随机整数
//
//
//示例 1：
//
//输入
//["Solution", "pick", "pick", "pick", "pick", "pick", "pick", "pick"]
//[[7, [2, 3, 5]], [], [], [], [], [], [], []]
//输出
//[null, 0, 4, 1, 6, 1, 0, 4]
//
//解释
//Solution solution = new Solution(7, [2, 3, 5]);
//solution.pick(); // 返回0，任何[0,1,4,6]的整数都可以。注意，对于每一个pick的调用，
//// 0、1、4和6的返回概率必须相等(即概率为1/4)。
//solution.pick(); // 返回 4
//solution.pick(); // 返回 1
//solution.pick(); // 返回 6
//solution.pick(); // 返回 1
//solution.pick(); // 返回 0
//solution.pick(); // 返回 4
//
//
//提示:
//
//1 <= n <= 109
//0 <= blacklist.length <= min(105, n - 1)
//0 <= blacklist[i] < n
//blacklist 中所有值都 不同
//pick 最多被调用 2 * 104 次

func main() {
	s := Constructor(7, []int{2, 3, 5})
	fmt.Println(s.Pick())
	fmt.Println(s.Pick())
	fmt.Println(s.Pick())
	fmt.Println(s.Pick())
	fmt.Println(s.Pick())
	fmt.Println(s.Pick())
	fmt.Println(s.Pick())
	fmt.Println(s.Pick())
}

type Solution struct {
	sz int
	m  map[int]int //key=索引，value=索引
}

func Constructor(n int, blacklist []int) Solution {
	m := map[int]int{}
	sz := n - len(blacklist)

	for _, v := range blacklist {
		m[v] = 999
	}

	last := n - 1

	for _, v := range blacklist {
		if v >= sz {
			continue
		}
		for m[last] > 0 {
			last--
		}
		m[v] = last
		last--
	}
	return Solution{sz, m}
}

func (this *Solution) Pick() int {
	idx := rand.Intn(this.sz)
	//命中黑名单
	if newIdx, ok := this.m[idx]; ok {
		return newIdx
	}
	//没有命中黑名单
	return idx
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n, blacklist);
 * param_1 := obj.Pick();
 */
