package leetcodeHot100

func maxDepth(root *TreeNode) int {
	//递归出口
	if root == nil {
		return 0
	} else {
		//只需要考虑一件事，找到左右子树的最大长度就可以了，然后再加上本层的长度
		return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	}
}
