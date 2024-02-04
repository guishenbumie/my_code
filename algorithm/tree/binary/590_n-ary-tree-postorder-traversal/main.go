package main

/* N 叉树的后序遍历 */

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

func postorder(root *Node) []int {
	ans := []int{}
	var traverse func(root *Node)
	traverse = func(root *Node) {
		if root == nil {
			return
		}
		for _, n := range root.Children {
			traverse(n)
		}
		ans = append(ans, root.Val)
	}
	traverse(root)
	return ans
}
