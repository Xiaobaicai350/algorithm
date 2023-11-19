package leetcodeHot100

func maxArea(height []int) int {
	//用来记录结果
	ans := 0
	//定义左指针
	left := 0
	//定义右指针
	right := len(height) - 1
	for left < right {
		//记录当前能装的最大容量 = 短板的高度*长度
		currentAns := min(height[left], height[right]) * (right - left)
		//记录一下最大值
		ans = max(ans, currentAns)
		//移动一下短板，如果是左边的板短，就移动左边，右边同理
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}
