package main

/* 最长有效括号 */

func main() {

}

func longestValidParentheses(s string) int {
	stack := []int{}
	dp := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			dp[i+1] = 0
		} else {
			if len(stack) > 0 {
				leftIndex := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				l := 1 + i - leftIndex + dp[leftIndex]
				dp[i+1] = l
			} else {
				dp[i+1] = 0
			}
		}
	}
	ans := 0
	for _, x := range dp {
		if x > ans {
			ans = x
		}
	}
	return ans
}
