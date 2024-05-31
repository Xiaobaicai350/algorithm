package leetcodeHot100

/*
输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
*/
func lengthOfLIS(nums []int) int {
	// dp数组表示的含义是： 以nums[i]为结尾的最长递增子序列的长度是dp[i]
	dp := make([]int, len(nums))
	//初始化返回结果
	res := 1
	//初始化dp数组，都为1，因为是以nums[i]为结尾，所以至少为1
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	//从第一个数字开始循环
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] { //给dp数组赋值
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
