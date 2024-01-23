package main

/* 搜索旋转排序数组 II */

func main() {

}

// 这题没什么意义，可以用二分，但遍历一次不是直接就出来了，强行出题呀
func search(nums []int, target int) bool {
	for _, x := range nums {
		if x == target {
			return true
		}
	}
	return false
}
