package main

/* 所有可能的路径 */

func main() {

}

//图回头再看看

//func allPathsSourceTarget(graph [][]int) [][]int {
//	res := make([][]int, 0)
//	p := make([]int, 0) //递归过程中经过的路径
//	var traverse func(graph [][]int, s int, path []int)
//	traverse = func(graph [][]int, s int, path []int) {
//		path = append(path, s)
//		if s == len(graph)-1 {
//			//到达终点
//			//temp := make([]int, len(path))
//			//copy(temp, path)
//			res = append(res, path)
//			path = path[:len(path)-1]
//			return
//		}
//		for _, v := range graph[s] {
//			traverse(graph, v, path)
//		}
//		path = path[:len(path)-1]
//	}
//	traverse(graph, 0, p)
//	return res
//}
