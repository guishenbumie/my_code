package main

import (
	"fmt"
	"sort"
)

/* 删除被覆盖区间 */

func main() {
	fmt.Println(removeCoveredIntervals([][]int{
		{1, 4},
		{3, 6},
		{2, 8},
	}))
}

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		if a[0] < b[0] {
			return true
		}
		if a[0] == b[0] {
			return a[1] > b[1]
		}
		return false
	})
	left := intervals[0][0]
	right := intervals[0][1]
	res := 0
	for i := 1; i < len(intervals); i++ {
		v := intervals[i]
		//情况一，找到覆盖区间
		if left <= v[0] && right >= v[1] {
			res++
		}
		//情况二，找到相交区间
		if right >= v[0] && right <= v[1] {
			left = v[0] //其实这一步可以省略，因为已经排序过了
			right = v[1]
		}
		//情况三，完全不相交，更新起点终点
		if right < v[0] {
			left = v[0]
			right = v[1]
		}
	}
	return len(intervals) - res
}
