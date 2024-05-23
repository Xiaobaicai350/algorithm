package leetcodeHot100

/**
 * 这道题首先应该确定递归的终止条件
 * 其次是判断是否最近公共祖先是他们其中的一个
 * 之后以左右子树作为根节点传入进去得到以root.xx为根节点，返回p和q的公共祖先
 * 之后就可以进行判断了
 */
// lowestCommonAncestor 用于寻找以root为根的树中p和q节点的公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//递归的终止条件
	if root == nil {
		return nil
	}
	//如果是根节点和左孩子相同或者是根节点和右孩子相同，直接返回根节点（这种情况下就是最近公共祖先是他们中的其中一个）
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	//以root.left为根节点，返回p和q的公共祖先
	left := lowestCommonAncestor(root.Left, p, q)
	//以root.right为根节点，返回p和q的公共祖先
	right := lowestCommonAncestor(root.Right, p, q)
	//如果两边都没有找到，说明这个最近公共祖先就是这个根节点
	if left != nil && right != nil {
		return root
	}
	if left == nil { //如果在左边没有找到，就说明最近公共祖先是右边这个
		return right
	} else {
		return left
	}
}
