package main

import "math"

/* 最小覆盖子串 */

// 和76题一样
func main() {

}

func minWindow(s string, t string) string {
	need, windows := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	start, length := 0, math.MaxInt32

	left, right := 0, 0
	valid := 0
	for right < len(s) {
		c := s[right]
		right++

		if need[c] > 0 {
			windows[c]++
			if need[c] == windows[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				length = right - left
				start = left
			}

			d := s[left]
			left++

			if need[d] > 0 {
				if need[d] == windows[d] {
					valid--
				}
				windows[d]--
			}
		}
	}

	if length == math.MaxInt32 {
		return ""
	}

	return s[start : start+length]
}
