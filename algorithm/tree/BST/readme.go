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

// 在BST中搜索元素
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	//去左子树搜
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	//去右子树搜
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	return root
}

// 在BST中插入一个数
func insert2BST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	//BST中一般不会插入已经存在的元素
	//if root.Val == val {}
	if root.Val < val {
		root.Right = insert2BST(root.Right, val)
	}
	if root.Val > val {
		root.Left = insert2BST(root.Left, val)
	}
	return root
}

// 获取最小值，BST最左边就是最小值
func getMin(n *TreeNode) *TreeNode {
	for n.Left != nil {
		n = n.Left
	}
	return n
}

// 在BST中删除一个数
func deleteNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		//要删除的节点至少有一个孩子为空，孩子替换为root节点
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		//有两个孩子
		//获取右子树最小的孩子
		min := getMin(root)
		//删除这个孩子
		root.Right = deleteNode(root.Right, val)
		//然后用这个孩子替换当前的root节点
		min.Left = root.Left
		min.Right = root.Right
		root = min
	} else if root.Val > val {
		root.Left = deleteNode(root.Left, val)
	} else if root.Val < val {
		root.Right = deleteNode(root.Right, val)
	}
	return root
}
