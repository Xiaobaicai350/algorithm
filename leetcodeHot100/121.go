package leetcodeHot100

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		//如果长度为0或者为1
		//为1的情况说明是当天买当天卖，利润为0
		//为0的情况说明是空数组，利润当然也为0
		return 0
	}
	//dp[i]的含义是当第i天卖出去，可以获得的最大利润是dp[i]
	dp := make([]int, len(prices))
	//初始化dp数组
	dp[0] = 0
	//初始化买入数据
	buy := prices[0]
	for i := 1; i < len(prices); i++ {
		//计算当天的利润
		profit := prices[i] - buy
		dp[i] = max(profit, dp[i-1])
		//更新买入的值
		buy = min(prices[i], buy)
	}
	return dp[len(prices)-1]
}
