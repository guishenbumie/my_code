package main

/* 验证回文串 II */

func main() {

}

func validPalindrome(s string) bool {
	lo, hi := 0, len(s)-1
	for lo < hi {
		if s[lo] == s[hi] { //如果左右都相等，同时往里缩
			lo++
			hi--
		} else { //如果左右不等，说明其中一个该删掉，然后剩下的如果是回文就满足条件了
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

//下面的时间复杂度O(n^2)，太高了
// func validPalindrome(s string) bool {
//     if valid(s) {
//         return true
//     }

//     for i:=0; i<len(s); i++ {
//         s1 := s[:i]
//         s2 := s[i+1:len(s)]
//         newStr := strings.Join([]string{s1, s2}, "")
//         if valid(newStr) {
//             return true
//         }
//     }
//     return false
// }

// func valid(s string) bool {
//     for i:=0; i<len(s)/2; i++ {
//         if s[i] != s[len(s)-i-1] {
//             return false
//         }
//     }
//     return true
// }
