package main

import "math"

/* 区间列表的交集 */

func main() {

}

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	ans := [][]int{}
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		a1 := firstList[i][0]
		a2 := firstList[i][1]
		b1 := secondList[j][0]
		b2 := secondList[j][1]

		if b2 >= a1 && a2 >= b1 {
			ans = append(ans, []int{
				int(math.Max(float64(a1), float64(b1))),
				int(math.Min(float64(a2), float64(b2))),
			})
		}
		if b2 < a2 {
			j++
		} else {
			i++
		}
	}
	return ans
}
