package leetcodeHot100

func climbStairs(n int) int {
	//如果n为1，说明只有一种方法
	if n == 1 {
		return 1
	}
	//创建dp数组
	//dp[i]的含义是：当楼梯层数为i时，有dp[i]种方法爬上去
	dp := make([]int, n+1)
	//初始化数组
	dp[1] = 1
	dp[2] = 2
	//这里注意是从3开始，直到n的情况
	for i := 3; i <= n; i++ {
		//递推公式就是前面两个dp的相加
		dp[i] = dp[i-1] + dp[i-2]
	}
	//直接返回最后一个
	return dp[n]
}
