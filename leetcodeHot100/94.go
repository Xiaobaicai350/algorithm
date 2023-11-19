package leetcodeHot100

var resArr []int

func inorderTraversal(root *TreeNode) []int {
	resArr = make([]int, 0)
	inorder(root)
	return resArr
}
func inorder(node *TreeNode) {
	if node == nil {
		return
	}
	inorder(node.Left)
	resArr = append(resArr, node.Val)
	inorder(node.Right)
}
