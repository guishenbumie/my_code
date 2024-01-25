package main

/* 招式拆解 I */

// 和3题一样
func main() {

}

func dismantlingAction(arr string) int {
	ans := 0
	window := map[byte]int{}

	left, right := 0, 0
	for right < len(arr) {
		c := arr[right]
		right++
		window[c]++

		for window[c] > 1 {
			d := arr[left]
			window[d]--
			left++
		}

		if right-left > ans {
			ans = right - left
		}
	}

	return ans
}
