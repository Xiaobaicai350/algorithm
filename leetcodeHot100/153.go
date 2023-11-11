package leetcodeHot100

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/solutions/698479/xun-zhao-xuan-zhuan-pai-xu-shu-zu-zhong-5irwp/?envType=study-plan-v2&envId=top-100-liked
func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] { //说明右边是正常的，可以直接忽略了
			right = mid
		} else { //说明右边是不正常的，左侧是正常的
			//表明nums[mid]是最小值左侧的元素，所以可以直接忽略左半部分，这里注意是+1
			left = mid + 1
		}
	}
	//到这里说明left==right了，所以返回哪个都可以了
	return nums[right]
}
