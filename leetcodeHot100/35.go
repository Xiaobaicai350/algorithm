package leetcodeHot100

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		//计算中间位置
		mid := l + (r-l)/2
		//如果中间位置的值大于等于目标值，说明在左半边
		if nums[mid] >= target {
			//在左半边需要改变右边界下标
			r = mid
		} else {
			//在右半边需要改变左边界下标
			l = mid + 1
		}
	}
	//这里返回l或者r都可以，因为退出上面循环的条件是l==r
	return r
}
