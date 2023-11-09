package leetcodeHot100

// 这道题没怎么懂，但是可以由递推公式推出来
func maxProduct(nums []int) int {
	ans := nums[0]
	dpMax := nums[0]
	dpMin := nums[0]
	for i := 1; i < len(nums); i++ {
		a := dpMax * nums[i]
		b := dpMin * nums[i]
		dpMax = max(nums[i], max(a, b))
		dpMin = min(nums[i], min(a, b))
		ans = max(ans, dpMax)
	}
	return ans
}
