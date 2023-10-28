package leetcode

/*
*
其实这段方法使用了前序遍历的方法
先交换中节点的左右节点指针，再进行翻转左子树和右子树
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	//中
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	//左
	invertTree(root.Left)
	//又
	invertTree(root.Right)
	return root
}
