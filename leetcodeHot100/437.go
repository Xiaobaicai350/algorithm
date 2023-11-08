package leetcodeHot100

// dfs的意思是 以root为根节点，可以找到res个符合题意的结果
func dfs(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	val := root.Val
	if val == targetSum {
		res++
	}
	res += dfs(root.Left, targetSum-val)
	res += dfs(root.Right, targetSum-val)
	return
}

func pathSum(root *TreeNode, targetSum int) int {
	//如果根节点都为空，说明这棵树就不存在，直接返回0
	if root == nil {
		return 0
	}
	//以根节点找和为targetSum
	res := dfs(root, targetSum)
	//往左右进行遍历
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}
