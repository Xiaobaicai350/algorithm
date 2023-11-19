package leetcodeHot100

func maxSubArray(nums []int) int {
	//dp数组的含义是：以nums[n]为结尾的情况下，最大子数组的和是dp[n]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	//记录最大值
	maxVal := dp[0]
	for i := 1; i < len(nums); i++ {
		//dp[i]是由dp[i-1]推导出来的
		if dp[i-1] >= 0 { //如果前面那个dp大于等于0的话，可以直接复用一下
			dp[i] = dp[i-1] + nums[i]
		} else { //如果是小于0的话，就没必要服用了，直接用nums[i]就好了
			dp[i] = nums[i]
		}
		maxVal = max(dp[i], maxVal)
	}
	return maxVal
}
