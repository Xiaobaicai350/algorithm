package leetcodeHot100

var resFTN []int

func findTargetNode(root *TreeNode, cnt int) int {
	inorderFTN(root)
	return resFTN[len(resFTN)-cnt]
}

func inorderFTN(node *TreeNode) {
	if node == nil {
		return
	}
	inorderFTN(node.Left)
	resFTN = append(resFTN, node.Val)
	inorderFTN(node.Right)
}
