package leetcodeHot100

/*
*
其实这段方法使用了前序遍历的方法
先交换根节点的左右节点指针，再进行翻转左子树和右子树
其实递归就可以看成一个过程，首先要确定终止条件，然后确定函数的功能，不如这里，只需要关注函数的功能是翻转根节点的左右子树就好了，之后会自动退出的
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	//根
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	//左
	invertTree(root.Left)
	//右
	invertTree(root.Right)
	return root
}
