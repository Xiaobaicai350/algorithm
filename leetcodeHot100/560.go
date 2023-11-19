package leetcodeHot100

func subarraySum(nums []int, k int) int {
	//用来表示从nums[0]...到nums[i]的前缀和
	sum := 0
	//记录总数
	count := 0
	//key是前缀和，value是前缀和出现了多少次
	m := map[int]int{}
	//前缀和为0的情况，也就是什么都没有的情况只有1次
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		//更新前缀和
		sum += nums[i]
		//如果存在前缀和为sum-k，就把count加上
		if _, ok := m[sum-k]; ok {
			count += m[sum-k] // 如果存在，将之前的出现次数加到 count 中
		}
		//这一步不能提前，必须在加完count之后写
		//将当前前缀和的出现次数加1
		m[sum] += 1
	}
	// 返回满足条件的子数组数量
	return count
}
