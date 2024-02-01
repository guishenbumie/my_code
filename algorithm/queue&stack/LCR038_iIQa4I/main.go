package main

/* 每日温度 */

// 和739题一样
func main() {

}

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	s := []int{}
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(s) > 0 && temperatures[s[len(s)-1]] <= temperatures[i] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			ans[i] = 0
		} else {
			ans[i] = s[len(s)-1] - i
		}
		s = append(s, i)
	}
	return ans
}
