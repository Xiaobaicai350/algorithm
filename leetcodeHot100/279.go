package leetcodeHot100

import "math"

func numSquares(n int) int {
	//创建数组切片
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minn = min(minn, dp[i-j*j])
		}
		dp[i] = minn + 1
	}
	return dp[n]
}
