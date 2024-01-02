package main

import "container/list"

func main() {

}

// 图节点的基本逻辑结构
type Vertex struct {
	id        int
	neighbors []Vertex
}

// 邻接表，存储x的所有相邻的节点
// 占用空间少，但无法快速判断两个节点是否相邻
var graph []list.List

// 邻接矩阵，matrix[x][y]记录x是否有一条指向y的记录；如果是加权有向图，bool就变成int
var matrix [][]bool

//// 图的遍历
//var visited []bool //记录被遍历过的节点
//var onPath []bool  //记录从起点到当前节点的路径
//func traverse(graph Graph, s int) {
//	if visited[s] {
//		return
//	}
//	//经过节点s，标记为已经遍历
//	visited[s] = true
//	//做选择：标记s在路径上
//	onPath[s] = true
//	for _, neighbor := range graph.neighbors[s] {
//		traverse(graph, neighbor)
//	}
//	//撤销选择：节点s离开路径
//	onPath[s] = false
//}
