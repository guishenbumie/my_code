package main

//func main() {
//	fmt.Println(isMatch("aa", "a"))
//	fmt.Println(isMatch("aa", "a*"))
//	fmt.Println(isMatch("ab", "*b"))
//}
//
//func isMatch(s string, p string) bool {
//	var dp func(c1 []byte, i int, c2 []byte, j int) bool
//	dp = func(c1 []byte, i int, c2 []byte, j int) bool {
//		m := len(c1)
//		n := len(c2)
//		if j == n {
//			return i == m
//		}
//		if i == m {
//			if (n-j)%2 == 1 {
//				return false
//			}
//			for ; j+1 < n; j += 2 {
//				if p[j+1] != '*' {
//					return false
//				}
//			}
//			return true
//		}
//	}
//
//	i, j := 0, 0
//	c1 := []byte(s)
//	c2 := []byte(p)
//	for i < len(c1) && j < len(c2) {
//		if c1[i] == c2[j] || p[j] == '.' { //匹配
//			if j < len(c2)-1 && c2[j+1] == '*' { //有*通配符，可以匹配0或者多次
//				return dp(c1, i, c2, j+2) || dp(c1, i+1, c2, j)
//			} else { //无*通配符，只能匹配一次
//				return dp(c1, i+1, c2, j+1)
//			}
//		} else { //不匹配
//			if j < len(c2)-1 && c2[j+1] == '*' { //有*通配符，只能匹配0次
//				return dp(c1, i, c2, j+2)
//			} else { //无*通配符，不能继续匹配
//				return false
//			}
//		}
//	}
//	return i == j
//}
