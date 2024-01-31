package main

/* 和为 K 的子数组 */

// 和560题一样
func main() {

}

func subarraySum(nums []int, k int) int {
	ans := 0
	prefix := 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		prefix += nums[i]
		if count, ok := m[prefix-k]; ok {
			ans += count
		}
		m[prefix] += 1
	}
	return ans
}
