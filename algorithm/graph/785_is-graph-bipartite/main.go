package main

/* 判断二分图 */

func main() {

}

// 双色问题，给你一幅「图」,请你用两种颜色将图中的所有顶点着色,且使得任意一条边的两个端点的颜色都不相同
func isBipartite(graph [][]int) bool {
	n := len(graph)
	ans := true
	color := make([]bool, n)   //记录图中节点的颜色，两种
	visited := make([]bool, n) //记录是否已经被访问过了

	var traverse func(x int)
	traverse = func(x int) {
		if !ans {
			return
		}
		visited[x] = true
		for _, y := range graph[x] {
			if !visited[y] { //相邻的节点y没有被访问过，则给y标记成和x不同的颜色
				color[y] = !color[x]
				traverse(y)
			} else { //如果y已经被访问过了，并且x和y的颜色一样，那就不是二分图
				if color[y] == color[x] {
					ans = false
				}
			}
		}
	}
	for x := 0; x < n; x++ {
		if !visited[x] {
			traverse(x)
		}
	}
	return ans
}
