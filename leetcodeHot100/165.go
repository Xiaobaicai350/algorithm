package leetcodeHot100

// compareVersion 比较两个版本号字符串 version1 和 version2 的大小
func compareVersion(version1, version2 string) int {
	n, m := len(version1), len(version2) // 获取两个字符串的长度
	i, j := 0, 0                         // 初始化两个指针，用于遍历字符串
	// 当任一字符串未遍历完时，继续比较
	for i < n || j < m {
		x := 0 // 用于存储 version1 的当前段数值
		// 遍历 version1，计算当前段数值，直到遇到点号或字符串结束
		for i < n && version1[i] != '.' {
			x = x*10 + int(version1[i]-'0') // 将字符转换为数字并累加到 x
			i++
		}
		i++    // 跳过点号，准备处理下一个段
		y := 0 // 用于存储 version2 的当前段数值
		// 遍历 version2，计算当前段数值，直到遇到点号或字符串结束
		for j < m && version2[j] != '.' {
			y = y*10 + int(version2[j]-'0') // 将字符转换为数字并累加到 y
			j++
		}
		j++ // 跳过点号，准备处理下一个段
		// 比较当前段的数值大小，如果 version1 大于 version2 则返回 1
		if x > y {
			return 1
		}
		// 如果 version1 小于 version2 则返回 -1
		if x < y {
			return -1
		}
	}
	// 如果所有段都相等，则返回 0
	return 0
}
