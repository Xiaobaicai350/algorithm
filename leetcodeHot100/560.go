package leetcodeHot100

// 这道题一看见就可以用前缀和来做，因为题目中说了：子数组是数组中元素的连续非空序列。（注意是连续）
/*
因为我们需要找到和为k的连续子数组的个数。通过计算前缀和，
我们可以将问题转化为求解两个前缀和之差等于k的情况。
假设数组的前缀和数组为sum，
其中sum[i]表示从数组起始位置到第i个位置的元素之和。那么对于任意的两个下标i和j（i < j），如果sum[j] - sum[i] = k，
即从第i个位置到第j个位置的元素之和等于k，那么说明从第i+1个位置到第j个位置的连续子数组的和为k。
通过遍历数组，计算每个位置的前缀和，并使用一个哈希表来存储每个前缀和出现的次数。
在遍历的过程中，我们检查是否存在sum[j] - k的前缀和，如果存在，说明从某个位置到当前位置的连续子数组的和为k
*/
func subarraySum(nums []int, k int) int {
	//目前的前缀和
	sum := 0
	//记录结果。表示一共有res种可能 和的值为k
	res := 0
	//和为key的时候，有value种可能
	m := make(map[int]int)
	//和为0的时候，有一种可能
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := m[sum-k]; ok {
			res += v
		}
		m[sum]++
	}
	return res
}
