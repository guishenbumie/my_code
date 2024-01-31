package main

/* 找出字符串中第一个匹配项的下标 */

func main() {

}

// 暴力求解
func strStr(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)

	if l2 == 0 {
		return 0
	}
	if l1 == 0 || l1 < l2 {
		return -1
	}

	for i := 0; i <= l1-l2; i++ {
		if haystack[i:i+l2] == needle {
			return i
		}
	}
	return -1
}

//KMP方式
