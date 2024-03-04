package main

/* 只出现一次的数字 II */

// 和137题一样
func main() {

}

func singleNumber(nums []int) int {
	m := map[int]int{}
	for _, x := range nums {
		m[x]++
	}
	for x, count := range m {
		if count == 1 {
			return x
		}
	}
	return 0
}
