package leetcodeHot100

func spiralOrder(matrix [][]int) []int {
	//行数
	m := len(matrix)
	//列数
	n := len(matrix[0])
	//定义上下左右边界
	up := 0
	down := m - 1
	left := 0
	right := n - 1
	//定义结果集
	ans := make([]int, 0)
	for {
		//1.首先从左往右遍历，123     5.从左往右遍历 5
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[up][i])
		}
		up++
		if up > down {
			break
		}
		//2.从上往下遍历 6，9
		for i := up; i <= down; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		//3.从又往左遍历，8，7
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[down][i])
		}
		down--
		if down < up {
			break
		}
		//4.从下往上遍历 4
		for i := down; i >= up; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return ans
}
