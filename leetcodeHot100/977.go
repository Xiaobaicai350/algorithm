package leetcodeHot100

import "math"

// 百度一面
func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	index := len(nums) - 1
	ints := make([]int, len(nums))
	for left <= right {
		zuo := math.Pow(float64(nums[left]), 2)
		you := math.Pow(float64(nums[right]), 2)
		if zuo > you {
			ints[index] = int(zuo)
			left++
		} else {
			ints[index] = int(you)
			right--
		}
		index--
	}
	return ints
}
