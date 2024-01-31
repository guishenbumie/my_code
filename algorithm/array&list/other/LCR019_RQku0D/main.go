package main

/* 验证回文串 II */

// 和680题一样
func main() {

}

func validPalindrome(s string) bool {
	lo, hi := 0, len(s)-1
	for lo < hi {
		if s[lo] == s[hi] {
			lo++
			hi--
		} else {
			flag1, flag2 := true, true
			for i, j := lo, hi-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}
			for i, j := lo+1, hi; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}
