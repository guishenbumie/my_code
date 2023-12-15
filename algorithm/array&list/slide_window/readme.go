package main

func main() {

}

// 滑动窗口算法框架
func slidingWindow(s, t string) {
	need, window := map[byte]int{}, map[byte]int{}
	_ = window
	for i := 0; i < len(t); i++ {
		c := t[i]
		need[c]++
	}

	left, right := 0, 0
	for right < len(s) {
		//c是将移入窗口的字符
		c := s[right]
		_ = c
		//右移窗口
		right++
		//进行窗口内数据一系列更新
		//...
		//fmt.Printf("window:[%d, %d)\n", left, right)

		//判断左侧窗口是否要收紧
		for /*条件*/ {
			//d是将移出窗口的字符
			d := s[left]
			_ = d
			//左移窗口
			left++
			//进行窗口内数据一系列更新
			//...
		}
	}
}
