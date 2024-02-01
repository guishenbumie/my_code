package main

import "fmt"

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

func main() {
	fmt.Println(removeDuplicateLetters("bcabc"))
	fmt.Println(removeDuplicateLetters("cbacdcbc"))
}

func removeDuplicateLetters(s string) string {
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
	//// 每个字符出现次数
	//var count [26]int
	//// 是否在栈中，存在为true
	//var exist [26]bool
	//// 单调栈
	//var stack []rune
	//
	//// 统计字符出现次数
	//// 注：c-'a'的结果是数字，即距离字母'a'有多远，比如'a'-'a'=0, 即就是a本身
	//for _, c := range s {
	//	count[c-'a']++
	//}
	//
	//for _, c := range s {
	//	// 栈中已有c，跳过
	//	if exist[c-'a'] {
	//		// 同时减少这个字符出现的次数
	//		count[c-'a']--
	//		continue
	//	}
	//
	//	// 出栈的核心判断要素：
	//	// 1. 栈里有元素
	//	// 2. 栈顶元素大于当前元素c
	//	// 3. 栈顶元素在后续出现
	//	for len(stack) > 0 && stack[len(stack)-1] > c && count[stack[len(stack)-1]-'a'] > 0 {
	//		// 进入这里说明栈顶元素大于当前元素，所以不符合字典序，要把栈顶元素出栈
	//		// 【重要】for 循环确保栈里面保存的都是比当前元素小的，因为大的都被pop掉了
	//
	//		// 标记为栈中不含栈顶元素
	//		exist[stack[len(stack)-1]-'a'] = false
	//		// 删除栈顶元素
	//		stack = stack[:len(stack)-1]
	//	}
	//
	//	// 添加新字符
	//	stack = append(stack, c)
	//	// 减少该字符出现次数
	//	count[c-'a']--
	//	// 标记栈中有该字符
	//	exist[c-'a'] = true
	//}
	//
	//return string(stack)
	//left := ['z' + 1]int{} // 相比创建一个长为 26 的数组，多开一点空间更方便
	//for _, c := range s {
	//	left[c]++ // 统计每个字母的出现次数
	//}
	//ans := []rune{}
	//inAns := ['z' + 1]bool{}
	//for _, c := range s {
	//	left[c]--
	//	if inAns[c] { // ans 中不能有重复字母
	//		continue
	//	}
	//	for len(ans) > 0 && c < ans[len(ans)-1] && left[ans[len(ans)-1]] > 0 {
	//		// 如果 c < x，且右边还有 x，那么可以把 x 去掉，因为后面可以重新把 x 加到 ans 中
	//		x := ans[len(ans)-1]
	//		ans = ans[:len(ans)-1]
	//		inAns[x] = false // 标记 x 不在 ans 中
	//	}
	//	ans = append(ans, c) // 把 c 加到 ans 的末尾
	//	inAns[c] = true      // 标记 c 在 ans 中
	//}
	//return string(ans)
}
