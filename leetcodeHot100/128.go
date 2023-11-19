package leetcodeHot100

import "sort"

func longestConsecutive(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}
	//对数组进行升序排列
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	currentLen := 1
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] { //如果前一个和当前的元素相同，就进行下一个的验证
			continue
		} else if nums[i] == nums[i-1]+1 { //如果刚好连续，则把当前长度+1，之后再更新max值
			currentLen++
			maxLen = max(maxLen, currentLen)
		} else { //否则就说明不连续，中断了，需要重制当前连续长度
			currentLen = 1
		}
	}
	return maxLen
}
