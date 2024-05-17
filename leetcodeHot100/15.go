package leetcodeHot100

import "sort"

/*
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
*/
func threeSum(nums []int) [][]int {
	//结果集
	var result [][]int
	//对数组排序
	sort.Ints(nums)
	// 排序之后如果第一个元素已经大于零，那么无论如何组合都不可能凑成三元组，直接返回结果就可以了
	// 这段话去掉也不影响结果，只是一个优化的方案
	if nums[0] > 0 {
		return result
	}
	for i := 0; i < len(nums); i++ {
		// 由于是排序后的，所以为了去除重复解，从第二个数和前一个数开始比较，若相同就跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//定义双指针，一个指向当前位置的后一个位置，一个指向数组最后的位置
		left := i + 1
		right := len(nums) - 1
		for right > left {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 { //如果大于0，说明大了，应该让右指针往左移动
				right--
			} else if sum < 0 { //如果小于0，说明小了，应该让左指针向右移动
				left++
			} else {
				//符合结果
				result = append(result, []int{nums[i], nums[left], nums[right]})
				//跳过重复的结果
				for right > left && nums[right] == nums[right-1] {
					right--
				}
				//跳过重复的结果
				for right > left && nums[left] == nums[left+1] {
					left++
				}
				//继续进行遍历，看剩下的数组还有没有合适的解
				right--
				left++
			}
		}
	}
	return result
}
