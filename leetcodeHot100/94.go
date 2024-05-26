package leetcodeHot100

func inorderTraversal(root *TreeNode) []int {
	resArr := make([]int, 0)
	inorder(root, &resArr)
	return resArr
}

// resArr现在是一个指向切片的指针
func inorder(node *TreeNode, resArr *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, resArr)
	*resArr = append(*resArr, node.Val)
	inorder(node.Right, resArr)
}
