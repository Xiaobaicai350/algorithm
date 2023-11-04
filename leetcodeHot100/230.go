package leetcodeHot100

func kthSmallest(root *TreeNode, k int) int {
	vec := traversal(root)
	return vec[k]
}
