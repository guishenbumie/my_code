package main

import "math"

/* 无重复字符的最长子串 */

// 和3题一样
func main() {

}

func lengthOfLongestSubstring(s string) int {
	windows := map[byte]int{}
	ans := 0
	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		right++
		windows[c]++

		for windows[c] > 1 {
			d := s[left]
			windows[d]--
			left++
		}

		ans = int(math.Max(float64(ans), float64(right-left)))
	}
	return ans
}
