package leetcodeHot100

func numIslands(grid [][]byte) int {
	//结果集
	res := 0
	//安全性校验，如果长度为0就直接返回有0个岛屿
	if len(grid) == 0 {
		return 0
	}
	//得到row和col
	row, col := len(grid), len(grid[0])

	var dfs func(grid [][]byte, r, c int)
	//dfs 用来把这个下标下的上下左右方位都修改成0，防止重复计数
	dfs = func(grid [][]byte, r, c int) {
		//如果访问的越界，或者访问的为0，就跳出去。ps:第一行检验越界，第二行检验是否为0
		if r < 0 || c < 0 || r >= row || c >= col ||
			grid[r][c] == '0' {
			return
		}
		//到这里说明没越界也不为0，为1
		grid[r][c] = '0'
		//遍历附近的四个方位是否也为1，有1就修改成0
		dfs(grid, r-1, c)
		dfs(grid, r+1, c)
		dfs(grid, r, c-1)
		dfs(grid, r, c+1)
	}

	//遍历整个二维数组
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == '1' {
				res++
				dfs(grid, r, c)
			}
		}
	}

	return res
}
