package leetcodeHot100

// https://www.bilibili.com/video/BV1aF411H7fn/?spm_id_from=333.337.search-card.all.click&vd_source=2259e5459a8cfd21bcf92bc46bf3beda
func searchRange(nums []int, target int) []int {
	firstIndex := searchFirst(nums, target)
	lastIndex := searchLast(nums, target)
	return []int{firstIndex, lastIndex}
}

func searchFirst(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			if mid == 0 || nums[mid-1] != nums[mid] {
				return mid
			} else {
				right = mid - 1
			}
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func searchLast(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			if mid == len(nums)-1 || nums[mid] != nums[mid+1] {
				return mid
			} else {
				left = mid + 1
			}
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
