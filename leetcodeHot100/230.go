package leetcodeHot100

var ints []int

func kthSmallest(root *TreeNode, k int) int {
	ints = make([]int, 0)
	inorderKth(root)
	return ints[k]
}
func inorderKth(root *TreeNode) {
	if root == nil {
		return
	}
	inorderKth(root.Left)
	ints = append(ints, root.Val)
	inorderKth(root.Right)
}
