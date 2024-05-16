package leetcodeHot100

func maxArea(height []int) int {
	//结果集，存储最大的面积
	ans := 0
	//定义左右指针，指向数组的两边的下标
	left := 0
	right := len(height) - 1
	//当左右指针相遇的时候会退出循环
	for left < right {
		if height[left] < height[right] { //如果右边的高度比较高
			//取短板作为自己的高
			ans = max(ans, height[left]*(right-left))
			//移动短板，这样的话才有可能会有更大的面积
			left++
		} else {
			ans = max(ans, height[right]*(right-left))
			right--
		}
	}
	//返回结果
	return ans
}
