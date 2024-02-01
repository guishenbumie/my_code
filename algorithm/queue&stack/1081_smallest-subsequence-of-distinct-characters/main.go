package main

/* 不同字符的最小子序列 */

// 和316题一样
func main() {

}

func smallestSubsequence(s string) string {
	//26个英文字母出现的次数
	count := make([]int, 26)
	for _, c := range s {
		count[c-'a']++
	}
	//字母是否已经在栈中
	exist := make([]bool, 26)
	//栈
	sta := []byte{}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if exist[c-'a'] {
			count[c-'a']--
			continue
		}

		for len(sta) > 0 && sta[len(sta)-1] > c && count[sta[len(sta)-1]-'a'] > 0 {
			exist[sta[len(sta)-1]-'a'] = false
			sta = sta[:len(sta)-1]
		}

		sta = append(sta, c)
		count[c-'a']--
		exist[c-'a'] = true
	}
	return string(sta)
}
