package main

/* 望远镜中最高的海拔 */

// 和239题一样
func main() {

}

func maxAltitude(heights []int, limit int) []int {
	ans := []int{}

	q := []int{}
	for i, x := range heights {
		for len(q) > 0 && heights[q[len(q)-1]] < x {
			q = q[:len(q)-1]
		}
		q = append(q, i)

		if i-q[0] >= limit {
			q = q[1:]
		}

		if i >= limit-1 {
			ans = append(ans, heights[q[0]])
		}

	}

	return ans
}
