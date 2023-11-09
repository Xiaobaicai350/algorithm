package leetcodeHot100

func lengthOfLIS(nums []int) int {
	// dp数组表示的含义是： 以nums[i]为结尾的最长递增子序列的长度是dp[i]
	dp := make([]int, len(nums))
	//初始化dp数组，都为1，因为是以nums[i]为结尾，所以至少为1
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	//初始化返回值
	res := dp[0]
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		//给结果赋值，取最大的dp[i]
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}
