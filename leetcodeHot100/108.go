package leetcodeHot100

func sortedArrayToBST(nums []int) *TreeNode {
	//终止条件，最后数组为空就可以返回了
	if len(nums) == 0 {
		return nil
	}
	idx := len(nums) / 2
	root := &TreeNode{
		Val: nums[idx],
	}
	root.Left = sortedArrayToBST(nums[:idx])
	root.Right = sortedArrayToBST(nums[idx+1:])
	return root
}
