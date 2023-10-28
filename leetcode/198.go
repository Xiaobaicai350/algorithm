package leetcode

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	//初始化dp数组
	//dp[i]表示前i间房屋能偷窃到的最多钱数
	dp := make([]int, len(nums))
	//初始化dp数组
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		/**
		1.偷窃第 k 间房屋，那么就不能偷窃第 k−1 间房屋，偷窃总金额为前 k−2 间房屋的最高总金额与第 k 间房屋的金额之和。
		2.不偷窃第 k 间房屋，偷窃总金额为前 k−1 间房屋的最高总金额。
		之后取这两者的最大值就好了
		*/
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}
