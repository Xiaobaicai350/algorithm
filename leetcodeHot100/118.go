package leetcodeHot100

/*
*
1
11
121
1331
*/
func generate(numRows int) [][]int {
	//给二维数组分配内存
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		//给一维数组分配内存，注意i是从0开始的，所以要+1
		res[i] = make([]int, i+1)
		//首先给首末
		res[i][0] = 1
		res[i][i] = 1
		//其他位置就是用该位置的上一个位置和上一个位置前一个位置进行得出来的
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}
	return res
}
