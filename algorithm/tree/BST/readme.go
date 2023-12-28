package main

func main() {

}

//二叉搜索树两个关键点：
//1.左小右大,即每个节点的左子树都比当前节点的值小,右子树都比当前节点的值大。
//2.中序遍历结果是有序的

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 常用框架
func BST(root *TreeNode, target int) {
	if root.Val == target {
		//找到目标
	}
	if root.Val < target {
		BST(root.Right, target)
	}
	if root.Val > target {
		BST(root.Left, target)
	}
}

// 是否是一个标准的BST树
func IsBST(root *TreeNode) bool {
	var f func(r, min, max *TreeNode) bool
	f = func(r, min, max *TreeNode) bool {
		if r == nil {
			return true
		}
		//必须满足左子树最大值<当前根节点的值<右子树最小值
		if min != nil && r.Val <= min.Val {
			return false
		}
		if max != nil && r.Val >= max.Val {
			return false
		}
		//递归，左子树最大值为当前节点值；右子树最小值为当前节点值
		return f(root.Left, min, root) && f(root.Right, root, max)
	}
	return f(root, nil, nil)
}
