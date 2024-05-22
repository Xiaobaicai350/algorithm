package leetcodeHot100

func maxProfit(prices []int) int {
	//dp[i]的含义是，当第i+1天卖出股票时，获得最大的利润是dp[i]
	dp := make([]int, len(prices))
	//第1天卖出股票，获得最大利润是0
	dp[0] = 0
	buy := prices[0]
	for i := 1; i < len(prices); i++ {
		//计算利润
		profit := prices[i] - buy
		dp[i] = max(profit, dp[i-1])
		//更新买入的值
		buy = min(prices[i], buy)
	}
	return dp[len(prices)-1]
}
