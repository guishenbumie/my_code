package main

import (
	"math"
	"sort"
)

/* 合并区间 */

func main() {

}

func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		last := ans[len(ans)-1]
		if curr[0] <= last[1] {
			last[1] = int(math.Max(float64(last[1]), float64(curr[1])))
		} else {
			ans = append(ans, curr)
		}
	}
	return ans
}
