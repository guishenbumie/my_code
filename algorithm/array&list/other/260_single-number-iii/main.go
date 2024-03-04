package main

/* 只出现一次的数字 III */

func main() {

}

func singleNumber(nums []int) []int {
	m := map[int]int{}
	for _, x := range nums {
		m[x]++
	}
	ans := []int{}
	for x, count := range m {
		if count == 1 {
			ans = append(ans, x)
		}
	}
	return ans
}
