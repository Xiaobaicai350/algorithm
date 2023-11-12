package leetcodeHot100

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	}
}
