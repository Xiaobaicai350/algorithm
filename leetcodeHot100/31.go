package leetcodeHot100

// https://bilibili.com/video/BV1qb4y1q7kQ/?spm_id_from=333.337.search-card.all.click&vd_source=2259e5459a8cfd21bcf92bc46bf3beda
func nextPermutation(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				BubbleSort(nums, i+1, len(nums))
				return
			}
		}
	}
	//如果到这里还没返回的话，说明就是3，2，1这种情况
	BubbleSort(nums, 0, len(nums))
	return
}
