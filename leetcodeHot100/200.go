package leetcodeHot100

func numIslands(grid [][]byte) int {
	//结果集
	res := 0
	//安全性校验，如果
	if len(grid) == 0 {
		return 0
	}

	nr, nc := len(grid), len(grid[0])

	var dfs func(grid [][]byte, r, c int)
	dfs = func(grid [][]byte, r, c int) {
		nr := len(grid)
		nc := len(grid[0])

		if r < 0 || c < 0 || r >= nr || c >= nc || grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0'
		dfs(grid, r-1, c)
		dfs(grid, r+1, c)
		dfs(grid, r, c-1)
		dfs(grid, r, c+1)
	}

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				res++
				dfs(grid, r, c)
			}
		}
	}

	return res
}
