package main

/* N 叉树的层序遍历 */

func main() {

}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	q := []*Node{root}
	for len(q) > 0 {
		sz := len(q)
		arr := make([]int, sz)
		for i := 0; i < sz; i++ {
			arr[i] = q[i].Val
			for _, node := range q[i].Children {
				q = append(q, node)
			}
		}
		q = q[sz:]
		ans = append(ans, append([]int(nil), arr...))
	}
	return ans
}
