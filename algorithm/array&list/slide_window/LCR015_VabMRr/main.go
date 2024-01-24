package main

/* 找到字符串中所有字母异位词 */

func main() {

}

// 和438题一样
func findAnagrams(s string, p string) []int {
	ans := []int{}

	need, windows := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	left, right := 0, 0
	valid := 0
	for right < len(s) {
		c := s[right]
		right++

		if need[c] > 0 {
			windows[c]++
			if windows[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(p) {
			if valid == len(need) {
				ans = append(ans, left)
			}

			d := s[left]
			left++

			if need[d] > 0 {
				if windows[d] == need[d] {
					valid--
				}
				windows[d]--
			}
		}
	}

	return ans
}
