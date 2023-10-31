package leetcodeHot100

func subarraySum(nums []int, k int) int {
	sum := 0   // 初始化当前子数组的和为0
	count := 0 // 记录满足条件的子数组数量
	//key是前缀和，value是前缀和出现了多少次
	m := map[int]int{} // 用于存储前缀和的出现次数的哈希表
	m[0] = 1           // 初始化前缀和为0的出现次数为1，用于处理从数组开始位置就满足条件的情况
	for i := 0; i < len(nums); i++ {
		sum += nums[i]             // 更新当前子数组的和
		if _, ok := m[sum-k]; ok { // 检查是否存在之前的前缀和等于 sum - k
			count += m[sum-k] // 如果存在，将之前的出现次数加到 count 中
		}
		m[sum] += 1 // 将当前前缀和的出现次数加1
	}
	return count // 返回满足条件的子数组数量
}
