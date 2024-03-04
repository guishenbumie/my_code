package main

/* 连续数组 */

func main() {

}

func findMaxLength(nums []int) int {
	m := map[int]int{0: -1} // key 来储存 cur 值, value 来储存当前 index
	count := 0
	ans := 0
	//由于以上碰1加一，碰0减一的操作，当0与1数量一致时(连续数组), 其连续数组的和为零
	//因此我们知道数组前面的 cur 值是什么，在到达该连续数组尾部时就不会变。
	//因此我们只需要检查哈希表中是否存在其相同的 cur 值即可
	for i, x := range nums {
		if x == 1 {
			count++
		} else {
			count--
		}
		if prevIdx, ok := m[count]; ok {
			if i-prevIdx > ans {
				ans = i - prevIdx
			}
		} else {
			m[count] = i
		}
	}
	return ans
}
