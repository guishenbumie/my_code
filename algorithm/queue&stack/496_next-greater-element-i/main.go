package main

/* 下一个更大元素 I */

func main() {

}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextGreaters := nextGreater(nums2)
	m := map[int]int{}
	for k, v := range nums2 {
		m[v] = nextGreaters[k]
	}
	ans := make([]int, len(nums1))
	for i, v := range nums1 {
		ans[i] = m[v]
	}
	return ans
}

func nextGreater(nums []int) []int {
	res := make([]int, len(nums)) //答案
	s := []int{}                  //栈
	//倒序往栈里放
	for i := len(nums) - 1; i >= 0; i-- {
		//判断个子高矮
		for len(s) > 0 && s[len(s)-1] <= nums[i] {
			s = s[:len(s)-1] // 矮个都丢弃
		}
		//nums[i]身后的下一个更大元素
		if len(s) == 0 {
			res[i] = -1
		} else {
			res[i] = s[len(s)-1]
		}
		s = append(s, nums[i])
	}
	return res
}
