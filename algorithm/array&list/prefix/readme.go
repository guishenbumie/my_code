package main

// 前缀和主要适用的场景是原始数组不会被修改的情况下,频繁查询某个区间的累加和。
func main() {

}

// 核心代码
type PrefixSum struct {
	prefix []int
}

func Constructor(nums []int) PrefixSum {
	//通过数组构造前缀和，前缀和数组比源数组多一个，第一个元素是0
	prefix := make([]int, len(nums)+1)
	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i-1] + nums[i-1]
	}
	return PrefixSum{prefix: prefix}
}

func (p *PrefixSum) query(i, j int) int {
	return p.prefix[j+1] - p.prefix[i]
}
