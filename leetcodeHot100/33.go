package leetcodeHot100

func search(nums []int, target int) int {
	//这道题用的是二分查找
	left := 0
	right := len(nums) - 1
	for left <= right { //当left==right的时候仍然有效
		middle := left + ((right - left) / 2) //等同于 (left + right) / 2 目的是防止溢出
		if target == nums[middle] {
			return middle
		} else if nums[left] <= nums[middle] { //如果满足这个条件，说明目前的左半边是有序的，可以进行二分查找(注意这里必须是<=,因为是向下取整)
			if target >= nums[left] && target < nums[middle] { //如果在左边这个范围中
				right = middle - 1
			} else { //如果不在左边这个范围
				left = middle + 1
			}
		} else { //说明右半边有序
			if target > nums[middle] && target <= nums[right] { //如果在右边这个范围中
				left = middle + 1
			} else { //如果不在右边这个范围
				right = middle - 1
			}
		}
	}
	return -1
}
