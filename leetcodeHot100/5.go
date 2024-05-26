package leetcodeHot100

// longestPalindrome 函数用于寻找字符串s中的最长回文子串
func longestPalindrome(s string) string {
	if s == "" {
		return "" // 如果字符串为空，直接返回空串
	}
	start, end := 0, 0 // 初始化最长回文子串的起始和结束位置
	for i := 0; i < len(s); i++ {
		// 对每个字符，分别以它为中心或者以它和它下一个字符为中心进行扩展
		left1, right1 := expandAroundCenter(s, i, i)   // 它为中心进行扩展
		left2, right2 := expandAroundCenter(s, i, i+1) // 以它和它下一个字符为中心进行扩展
		// 更新这两个结果中最大的end和start
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	// 返回最长回文子串
	return s[start : end+1]
}

// expandAroundCenter函数用于从中心向外扩展，寻找以left和right为中心的最长回文子串
func expandAroundCenter(s string, left, right int) (int, int) {
	for {
		//当到达边界或者不等于的时候就退出循环
		if left < 0 || right >= len(s) || s[left] != s[right] {
			break
		}
		left--
		right++
	}
	// 这时候退出循环了，可以返回边界了
	return left + 1, right - 1
}
