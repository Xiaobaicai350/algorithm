package leetcodeHot100

var intsIsValid []int

func isValidBST(root *TreeNode) bool {
	intsIsValid = make([]int, 0)
	inorderIsValid(root)
	//之后验证数组是否是有序的就可以了
	for i := 1; i < len(intsIsValid); i++ {
		if intsIsValid[i] <= intsIsValid[i-1] {
			return false
		}
	}
	return true
}
func inorderIsValid(root *TreeNode) {
	if root == nil {
		return
	}
	inorderIsValid(root.Left)
	intsIsValid = append(intsIsValid, root.Val)
	inorderIsValid(root.Right)
}
