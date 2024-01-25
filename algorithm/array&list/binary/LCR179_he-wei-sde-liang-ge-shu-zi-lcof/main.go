package main

/* 查找总价格为目标值的两个商品 */

// 和167题一样
func main() {

}

func twoSum(price []int, target int) []int {
	for i := 0; i < len(price); i++ {
		left := i + 1
		right := len(price) - 1
		for left <= right {
			mid := left + (right-left)/2
			x := target - price[i]
			if price[mid] < x {
				left = mid + 1
			} else if price[mid] > x {
				right = mid - 1
			} else {
				return []int{price[i], price[mid]}
			}
		}
	}
	return []int{-1, -1}
}
