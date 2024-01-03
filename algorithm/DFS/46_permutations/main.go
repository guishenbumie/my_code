package main

import (
	"fmt"
)

/* 全排列 */

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//示例 2：
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//示例 3：
//
//输入：nums = [1]
//输出：[[1]]
//
//
//提示：
//
//1 <= nums.length <= 6
//-10 <= nums[i] <= 10
//nums 中的所有整数 互不相同

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

//var res *list.List
//
//func permute(nums []int) [][]int {
//	res = list.New()
//	track := list.New()
//	used := make([]bool, len(nums))
//	backtrack(nums, track, used)
//
//	result := make([][]int, 0)
//	for el := res.Front(); el != nil; el = el.Next() {
//		l := el.Value.(*list.List)
//		arr := make([]int, 0)
//		for e := l.Front(); e != nil; e = e.Next() {
//			num := e.Value.(int)
//			arr = append(arr, num)
//		}
//		result = append(result, arr)
//	}
//	return result
//}
//
//func backtrack(nums []int, track *list.List, used []bool) {
//	if track.Len() == len(nums) {
//		res.PushBack(track)
//		return
//	}
//	for i := 0; i < len(nums); i++ {
//		if used[i] {
//			continue
//		}
//		track.PushBack(nums[i])
//		used[i] = true
//		backtrack(nums, track, used)
//		track.Remove(track.Back())
//		used[i] = false
//	}
//}

func permute(nums []int) [][]int {
	//ans := make([][]int, 0)
	//n := len(nums)
	//path := make([]int, n)    //记录路径
	//onPath := make([]bool, n) //避免重复使用
	//var dfs func(int)
	//dfs = func(i int) {
	//	//触发结束条件
	//	if i == n {
	//		ans = append(ans, append([]int(nil), path...))
	//		return
	//	}
	//	for j, on := range onPath {
	//		//排除不合法的选择，对已经在path中出现过的了，排除
	//		if on {
	//			continue
	//		}
	//		//做选择
	//		path[i] = nums[j]
	//		onPath[j] = true
	//		//进入下一层决策
	//		dfs(i + 1)
	//		//取消选择
	//		onPath[j] = false
	//	}
	//}
	//dfs(0)
	//return ans
	ans := make([][]int, 0)
	path := make([]int, 0)
	onPath := make([]bool, len(nums))
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(nums) { //遍历到最后一层
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j := 0; j < len(nums); j++ {
			if onPath[j] {
				continue
			}
			path = append(path, nums[j])
			onPath[j] = true
			dfs(i + 1)
			path = path[:len(path)-1]
			onPath[j] = false
		}
	}
	dfs(0)
	return ans
}

// 如果不是返回全排列，而是返回元素个数为k的排列
func permute1(nums []int, k int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	onPath := make([]bool, len(nums))
	var dfs func(i int)
	dfs = func(i int) {
		if k == len(path) { //这里变为path长度和指定k相同
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j := 0; j < len(nums); j++ {
			if onPath[j] {
				continue
			}
			path = append(path, nums[j])
			onPath[j] = true
			dfs(i + 1)
			path = path[:len(path)-1]
			onPath[j] = false
		}
	}
	dfs(0)
	return ans
}
