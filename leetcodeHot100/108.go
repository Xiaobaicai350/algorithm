package leetcodeHot100

// 将有序数组转换成高度平衡的二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	//终止条件，最后数组为空就可以返回了
	if len(nums) == 0 {
		return nil
	}
	//找到中间位置
	idx := len(nums) / 2
	//得到TreeNode对象的指针
	root := &TreeNode{
		Val: nums[idx],
	}
	//注意下面两个逻辑都跳过了idx，因为都已经构造过idx了
	//传入nums中[0,idx)的值
	root.Left = sortedArrayToBST(nums[:idx])
	//传入nums中(idx,end]
	root.Right = sortedArrayToBST(nums[idx+1:])
	return root
}
