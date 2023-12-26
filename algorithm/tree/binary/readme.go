package main

func main() {

}

// 二叉树框架
type TreeNode struct {
	left  *TreeNode
	right *TreeNode
}

func Traverse(root *TreeNode) {
	//前序遍历
	Traverse(root.left)
	//中序遍历
	Traverse(root.right)
	//后序遍历
}
