package leetcodeHot100

// https://bilibili.com/video/BV14s4y1U77K/?spm_id_from=333.337.search-card.all.click&vd_source=2259e5459a8cfd21bcf92bc46bf3beda
func flatten(root *TreeNode) {
	for root != nil {
		//如果左子树为nil，就直接遍历下一个右子树
		if root.Left == nil {
			root = root.Right
		} else { //到这里说明左子树不为空，说明要把左子树进行翻转
			//保存左子树的地址
			pre := root.Left
			//找到左子树的最后一个右子树
			for pre.Right != nil {
				pre = pre.Right
			}
			//将最后一个右子树的右子树保存为根节点的右子树
			pre.Right = root.Right
			//将根节点的右子树替换成根节点的左子树
			root.Right = root.Left
			//将左子树制空
			root.Left = nil
			//继续向下遍历
			root = root.Right
		}
	}
}
