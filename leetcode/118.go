package leetcode

func generate(numRows int) [][]int {
	//给二维的定义长度
	ans := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		//给一维的定义长度
		ans[i] = make([]int, i+1)
		//首末位置都为1
		ans[i][0] = 1
		ans[i][i] = 1
		//其他位置单独处理
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}
