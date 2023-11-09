package leetcodeHot100

// 分割等和子集 动态规划
// 时间复杂度O(n^2) 空间复杂度O(n)
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 如果 nums 的总和为奇数则不可能平分成两个子集
	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	//dp数组表示的含义是：dp[i]表示背包的容量是i的时候，放进物品的最大重量是dp[i]
	dp := make([]int, target+1)

	for _, num := range nums {
		//这里一定得是倒叙遍历才能保证只能放一次
		for j := target; j >= num; j-- {
			if dp[j] < dp[j-num]+num {
				dp[j] = dp[j-num] + num
			}
		}
	}
	return dp[target] == target
}
