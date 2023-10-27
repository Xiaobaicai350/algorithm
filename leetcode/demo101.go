package leetcode

func defs(left *TreeNode, right *TreeNode) bool {
	//如果根节点的左右都为空，说明是合理情况
	if left == nil && right == nil {
		return true
	}
	//走到这里说明左右总会有一个不为空，如果是左右有一个为空就不行
	if left == nil || right == nil {
		return false
	}
	//到这里说明左右都不为空，可以比较左右是否相等了
	if left.Val != right.Val {
		//如果左右的值不相等，也是不符合条件的
		return false
	}
	//递归遍历外侧树和内侧树，他们俩必须也得都是true
	return defs(left.Left, right.Right) && defs(left.Right, right.Left)
}
func isSymmetric(root *TreeNode) bool {
	return defs(root.Left, root.Right)
}
