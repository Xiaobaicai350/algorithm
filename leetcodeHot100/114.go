package leetcodeHot100

// https://bilibili.com/video/BV14s4y1U77K/?spm_id_from=333.337.search-card.all.click&vd_source=2259e5459a8cfd21bcf92bc46bf3beda
// 该算法的步骤是：
// 1.当前节点不为nil的时候，就往下遍历
// 2.如果左节点为空，就往下遍历，找下一个左节点不为空的节点
// 3.当找到左节点不为空的时候，就找当前节点的左子树的最后一个右子树（这里是为了拼接当前节点的右子树）
// 4.之后就是拼接，然后把当前节点的左子树转换成右子树
// 5.把左子树置空，开始遍历下一个右子树，进行循环，直到为空
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
