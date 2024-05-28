package leetcodeHot100

func binSearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	mid := 0
	for left <= right {
		mid = (left + right) >> 1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
