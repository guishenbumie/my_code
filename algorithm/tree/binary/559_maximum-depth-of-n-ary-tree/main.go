package main

/* N 叉树的最大深度 */

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

func maxDepth(root *Node) int {
	dep := 0
	ans := 0
	var traverse func(root *Node)
	traverse = func(root *Node) {
		if root == nil {
			return
		}
		dep++
		if dep > ans {
			ans = dep
		}
		for _, n := range root.Children {
			traverse(n)
		}
		dep--
	}
	traverse(root)
	return ans
}
