package leetcodeHot100

var pos = make(map[int]int)

// 前序遍历
func buildTree(preorder []int, inorder []int) *TreeNode {
	//得到
	n := len(preorder)
	for i := 0; i < n; i++ {
		pos[inorder[i]] = i
	}
	return build(preorder, inorder, 0, n-1, 0, n-1)
}

func build(preorder []int, inorder []int, pl int, pr int, il int, ir int) *TreeNode {
	if pl > pr || il > ir {
		return nil
	}
	k := pos[preorder[pl]] - il
	root := &TreeNode{Val: preorder[pl]}
	root.Left = build(preorder, inorder, pl+1, pl+k, il, il+k-1)
	root.Right = build(preorder, inorder, pl+k+1, pr, il+k+1, ir)
	return root
}
