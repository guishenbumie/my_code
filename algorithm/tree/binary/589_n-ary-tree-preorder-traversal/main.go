package main

/* N 叉树的前序遍历 */

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

func preorder(root *Node) []int {
	ans := []int{}
	var traverse func(root *Node)
	traverse = func(root *Node) {
		if root == nil {
			return
		}
		ans = append(ans, root.Val)
		for _, n := range root.Children {
			traverse(n)
		}
	}
	traverse(root)
	return ans
}
