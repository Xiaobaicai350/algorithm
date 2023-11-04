package leetcodeHot100

func isValidBST(root *TreeNode) bool {
	//转换成数组
	vec := traversal(root)
	//之后验证数组是否是有序的就可以了
	for i := 1; i < len(vec); i++ {
		if vec[i] <= vec[i-1] {
			return false
		}
	}
	return true
}

func traversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	left := traversal(root.Left)
	right := traversal(root.Right)

	return append(append(left, root.Val), right...)
}
