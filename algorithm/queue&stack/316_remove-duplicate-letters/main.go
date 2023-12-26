package main

/* 去除重复字母 */

//给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。
//
//
//
//示例 1：
//
//输入：s = "bcabc"
//输出："abc"
//示例 2：
//
//输入：s = "cbacdcbc"
//输出："acdb"
//
//
//提示：
//
//1 <= s.length <= 104
//s 由小写英文字母组成

//func main() {
//	fmt.Println(removeDuplicateLetters("bcabc"))
//	fmt.Println(removeDuplicateLetters("cbacdcbc"))
//}
//
//func removeDuplicateLetters(s string) string {
//	left := ['z' + 1]int{} // 相比创建一个长为 26 的数组，多开一点空间更方便
//	for _, c := range s {
//		left[c]++ // 统计每个字母的出现次数
//	}
//	ans := []rune{}
//	inAns := ['z' + 1]bool{}
//	for _, c := range s {
//		left[c]--
//		if inAns[c] { // ans 中不能有重复字母
//			continue
//		}
//		for len(ans) > 0 && c < ans[len(ans)-1] && left[ans[len(ans)-1]] > 0 {
//			// 如果 c < x，且右边还有 x，那么可以把 x 去掉，因为后面可以重新把 x 加到 ans 中
//			x := ans[len(ans)-1]
//			ans = ans[:len(ans)-1]
//			inAns[x] = false // 标记 x 不在 ans 中
//		}
//		ans = append(ans, c) // 把 c 加到 ans 的末尾
//		inAns[c] = true      // 标记 c 在 ans 中
//	}
//	return string(ans)
//}
