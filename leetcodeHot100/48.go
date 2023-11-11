package leetcodeHot100

// 本题有很强的技巧性
// 顺时针旋转90度=上坐到下右分割线对称交换+竖中对称交换
func rotate(matrix [][]int) {
	//得到长度
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}
