package main

/* 字符串的排列 */

func main() {

}

// 和242题一样
func checkInclusion(s1 string, s2 string) bool {
	need, windows := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}

	valid := 0
	left, right := 0, 0

	for right < len(s2) {
		c := s2[right]
		right++

		if need[c] > 0 {
			windows[c]++
			if windows[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}

			d := s2[left]
			left++

			if need[d] > 0 {
				if windows[d] == need[d] {
					valid--
				}
				windows[d]--
			}
		}
	}

	return false
}
