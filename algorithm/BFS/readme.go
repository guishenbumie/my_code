package main

func main() {

}

//BFS，广度优先搜索算法，"湖面丢进一块石头激起层层涟漪"
//使用队列实现
//常用于找单一的最短路线，它的特点是 "搜到就是最优解"

//BFS的本质就是让你在一幅「图」中找到从起点start到终点 target的最近距离

//// 框架，计算从start到target的最近距离
//func BFS(start, target Node) int {
//	q := list.New() //用数组也行
//	visited := map[Node]struct{}{}
//
//	//起点加入队列
//	q.PushBack(start)
//	visited[start] = struct{}{}
//
//	var step int //步数
//	for q.Len() > 0 {
//		sz := q.Len()
//		//从当前节点向四周扩散
//		for i := 0; i < sz; i++ {
//			curr := q.Remove(q.Front()).(Node)
//			//是否到达终点
//			if curr == target {
//				return step
//			}
//			//将相邻的节点加入队列
//			for x := range curr.adj() { //伪代码，就是遍历的周围的节点
//				if _, ok := visited[x]; !ok {
//					q.PushBack(x)
//					visited[x] = struct{}{}
//				}
//			}
//		}
//		step++ //在这里进行步数+1
//	}
//}
