package leetcodeHot100

import "math"

// reverse 函数接收一个整数 x 并返回其反转后的整数。 如果反转后的整数超出了32位有符号整数的范围，则返回0。
func reverse7(x int) int {
	// 初始化结果变量 res 为 0
	res := 0
	// 当 x 不为0时，执行循环
	for x != 0 {
		// 如果 res 小于 math.MinInt32/10 (-214748364)或 res 大于 math.MaxInt32/10，
		// 则说明在下一步操作中 res 会超出32位有符号整数的范围。
		// 此时直接返回0。
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		// 获取 x 的个位数
		digit := x % 10
		// 将 x 除以10，用于获取下一位数
		x /= 10
		// 将 res 乘以10并加上个位数 digit，实现反转
		res = res*10 + digit
	}
	// 返回反转后的整数 res
	return res
}
