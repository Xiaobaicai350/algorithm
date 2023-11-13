package leetcodeHot100

// https://www.bilibili.com/video/BV1xA411g7C2/?spm_id_from=333.337.search-card.all.click&vd_source=2259e5459a8cfd21bcf92bc46bf3beda
func coinChange(coins []int, amount int) int {
	//dp数组代表的是：当钱数为i的时候，最少需要的钱数是dp[i]
	//所以数组的长度就是amount+1了
	dp := make([]int, amount+1)
	//金额0的最优解dp[0]=0
	dp[0] = 0
	//初始化数组其他值为-1
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}
	//变量从1循环到amount，依次计算金额1到amount的最优解
	for i := 1; i <= amount; i++ {
		//对于金额i，使用j遍历面值coins数组
		for j := 0; j < len(coins); j++ {
			//所有小于等于i的面值coin是[j],如果金额i-coins[j]有最优解
			if coins[j] <= i && dp[i-coins[j]] != -1 {
				//如果当前金额还未计算或者dp[i]比正在计算的最优解大
				if dp[i] == -1 || dp[i] > dp[i-coins[j]]+1 {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}
	return dp[amount]
}
