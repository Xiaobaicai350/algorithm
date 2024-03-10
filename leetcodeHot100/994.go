package leetcodeHot100

/*
*
值 0 代表空单元格；
值 1 代表新鲜橘子；
值 2 代表腐烂的橘子
*/
func orangesRotting(grid [][]int) int {
	good := 0         // 好橘子的数量
	bad := [][2]int{} // 坏橘子的位置

	// 遍历所有橘子，计算 好橘子的数量 和 坏橘子的位置
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			switch grid[i][j] {
			case 1: //如果值是1说明是新鲜橘子，就给计数加一
				good++
			case 2: //如果是2说明是腐烂的橘子，加入到数组中
				bad = append(bad, [2]int{i, j})
			}
		}
	}

	cnt := 0 // 记录扩散了多少次，相当于结果集

	// 使用 BFS 模拟橘子腐烂并扩散的过程
	for len(bad) > 0 {
		next := [][2]int{} // 记录本次腐烂的橘子

		// 开始扩散（遍历所有坏橘子）
		for len(bad) > 0 {
			//当前位置腐烂的橘子
			cur := bad[0]
			//剔除腐烂过后去橘子，防止重复遍历
			bad = bad[1:]
			i, j := cur[0], cur[1] // 拿到坏橘子的位置

			// 判断坏橘子是否在区域内
			if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
				//如果不在就跳过
				continue
			}

			// 往上下左右四个位置扩散（逻辑相同）
			if i-1 >= 0 && grid[i-1][j] == 1 {
				good--
				grid[i-1][j] = 2
				next = append(next, [2]int{i - 1, j})
			}
			if i+1 < len(grid) && grid[i+1][j] == 1 {
				good--
				grid[i+1][j] = 2
				next = append(next, [2]int{i + 1, j})
			}
			if j-1 >= 0 && grid[i][j-1] == 1 {
				good--
				grid[i][j-1] = 2
				next = append(next, [2]int{i, j - 1})
			}
			if j+1 < len(grid[0]) && grid[i][j+1] == 1 {
				good--
				grid[i][j+1] = 2
				next = append(next, [2]int{i, j + 1})
			}
		}

		// 本次腐烂的橘子数量大于0时，扩散次数+1
		if len(next) > 0 {
			cnt++
		}

		// 本次腐烂的橘子即为一下次扩散的开始
		bad = next
	}
	//退出循环之后，判断好橘子数量，如果还大于0，说明不可能全部腐烂
	// 有橘子扩散不到，返回 -1
	if good > 0 {
		return -1
	}

	return cnt
}
