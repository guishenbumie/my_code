package main

/* 寻找数组的中心下标 */

func main() {

}

func pivotIndex(nums []int) int {
	ans := -1
	preNums := make([]int, len(nums)+1)
	preNums[0] = 0
	for i := 1; i <= len(nums); i++ {
		preNums[i] = preNums[i-1] + nums[i-1]
	}
	for i := 0; i < len(nums); i++ {
		left := preNums[i]
		right := preNums[len(preNums)-1] - nums[i] - left
		if left == right {
			ans = i
			break
		}
	}
	return ans
}
