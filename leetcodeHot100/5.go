package leetcodeHot100

func longestPalindrome(s string) string {
	//创建dp数组，dp[i][j]的含义是：字符串从第i位下标到第j位下标是否是回文串
	dp := make([][]bool, len(s))
	//初始化结果(最小的回文就是单个字符)
	//也就是初始化字符串的第一位为结果
	result := s[0:1]
	for i := 0; i < len(s); i++ {
		//给一维数组分配空间
		dp[i] = make([]bool, len(s))
		// 给对角线上都初始化为true
		dp[i][i] = true // 根据case 1 初始数据
	}
	for length := 2; length <= len(s); length++ { //长度固定，不断移动起点
		for start := 0; start < len(s)-length+1; start++ { //长度固定，不断移动起点
			end := start + length - 1
			if s[start] != s[end] { //首尾不同则不可能为回文
				continue
			} else if length < 3 {
				//长度为2的字符串首尾相同就是回文
				dp[start][end] = true
			} else {
				dp[start][end] = dp[start+1][end-1] //状态转移
			}
			if dp[start][end] && (end-start+1) > len(result) { //记录最大值
				result = s[start : end+1]
			}
		}
	}
	return result
}
