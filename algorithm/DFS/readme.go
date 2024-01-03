package main

func main() {

}

//回溯算法，实际上就是一个决策树的遍历过程
//框架
//result = []
//def backtrack(路径,选择列表):
//	if 满足结束条件
//		result.add(路径)
//		return
//
//	for 选择 in 选择列表:
//		做选择
//		backtrack(路径,选择列表)
//		撤销选择

// //几种变化（组合和子集其实基本是等价的，排列的题要多一个数组来记录是否used）
// 1. 子集（元素无重不可复选）——78题
// 2. 组合（元素无重不可复选）——77题
// 3. 排列（元素无重不可复选）——46题
// 4. 子集/组合（元素可重不可复选）——90题&40题
// 5. 排列（元素可重不可复选）——47题
// 6. 子集/组合（元素无重可复选）——39题
// 7. 排列（元素无重可复选）——没有对应的题见下面
// 例如nums=[1,2,3]，可以复选共有多少种排列
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	//used := make([]bool, len(nums))
	var dfs func(i int)
	dfs = func(i int) {
		if len(path) == len(nums) {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j := 0; j < len(nums); j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}

// 框架
var path []int
var used []bool

// 一：元素无重不可复选，即nums中的元素都是唯一的,每个元素最多只能被使用一次
// 子集/组合框架
//	func backTrack(nums []int, start int) {
//		for i := start; i < len(nums); i++ {
//			//做选择
//			path = append(path, nums[i])
//			//递归回溯
//			backTrack(nums, i+1)
//			//撤销选择
//			path = path[:len(path)-1]
//		}
//	}
// 排列框架
//func backTrack(nums []int) {
//	for i := 0; i < len(nums); i++ {
//		//剪枝逻辑
//		if used[i] {
//			continue
//		}
//		//做选择
//		path = append(path, nums[i])
//		used[i] = true
//		//递归回溯
//		backTrack(nums)
//		//撤销选择
//		path = path[:len(path)-1]
//		used[i] = false
//	}
//}

// 二：元素重复不可复选，即nums中的元素可以存在重复,每个元素最多只能被使用一次
// 子集/组合框架
// sort.Ints(nums) //先排序
//
//	func backTrack(nums []int, start int) {
//		for i := start; i < len(nums); i++ {
//			//剪枝逻辑，跳过值相同的相邻树枝
//			if i > start && nums[i] == nums[i-1] {
//				continue
//			}
//			//做选择
//			path = append(path, nums[i])
//			//递归回溯
//			backTrack(nums, i+1)
//			//撤销选择
//			path = path[:len(path)-1]
//		}
//	}
//
// 排列框架
// sort.Ints(nums) //先排序
//
//	func backTrack(nums []int) {
//		for i := 0; i < len(nums); i++ {
//			//剪枝逻辑
//			if used[i] {
//				continue
//			}
//			//剪枝逻辑，固定相同的元素在排列中的位置
//			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
//				continue
//			}
//			//做选择
//			path = append(path, nums[i])
//			used[i] = true
//			//递归回溯
//			backTrack(nums)
//			//撤销选择
//			path = path[:len(path)-1]
//			used[i] = false
//		}
//	}
//
// 三：元素无重可复选，即nums中的元素都是唯一的,每个元素可以被使用若干次
// 子集/组合框架
//
//	func backTrack(nums []int, start int) {
//		for i := start; i < len(nums); i++ {
//			//做选择
//			path = append(path, nums[i])
//			//递归回溯
//			backTrack(nums, i)
//			//撤销选择
//			path = path[:len(path)-1]
//		}
//	}
//
// 排列框架
//func backTrack(nums []int) {
//	for i := 0; i < len(nums); i++ {
//		//做选择
//		path = append(path, nums[i])
//		//递归回溯
//		backTrack(nums)
//		//撤销选择
//		path = path[:len(path)-1]
//	}
//}
