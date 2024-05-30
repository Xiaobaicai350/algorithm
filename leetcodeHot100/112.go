package leetcodeHot100

// 定义hasPathSum函数，判断从根节点root到叶子节点的路径上是否存在节点值的和等于sum的情况
func hasPathSum(root *TreeNode, sum int) bool {
	// 如果根节点为空，直接返回false
	if root == nil {
		return false
	}
	// 如果当前节点是叶子节点，检查其值是否等于sum
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	// 递归检查左子树和右子树是否存在路径总和为sum减去当前节点值的路径
	// 使用逻辑或运算符，如果左子树或右子树有一条路径满足条件，就返回true
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
