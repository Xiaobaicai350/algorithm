package leetcodeHot100

// https://www.bilibili.com/video/BV1Ym4y1X72Q/?spm_id_from=333.337.search-card.all.click&vd_source=2259e5459a8cfd21bcf92bc46bf3beda
func search(nums []int, target int) int {
	//这道题就是用两个指针分别指向收尾，进行二分
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[left] <= nums[mid] { //如果左边界比中间的小的话
			//说明左边是没翻转的，也就是说左边是连续递增的
			if nums[left] <= target && target < nums[mid] { //说明target就在左边范围内，就可以进行二分了
				right = mid - 1
			} else { //否则就没在，说明在右边
				left = mid + 1
			}
		} else {
			//说明右边是没翻转的，也就是说右边是连续递增的
			if nums[mid] < target && target <= nums[right] { //说明target在右边范围内
				left = mid + 1
			} else { //否则就在左边
				right = mid - 1
			}
		}
	}
	return -1
}
