package main

/* 反转字符串 */

func main() {

}

func reverseString(s []byte) {
	var temp byte
	for i := 0; 2*i+1 < len(s); i++ {
		temp = s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = temp
	}
}
