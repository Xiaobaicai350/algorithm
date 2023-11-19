package leetcodeHot100

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	// 新建结果集
	var res [][]int
	// 按照左边界进行升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 初始化当前的左边界
	left := intervals[0][0]
	// 初始化当前的右边界
	right := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		// 如果下一个的左边界大于当前的右边界，说明没有重叠
		if intervals[i][0] > right {
			// 加入区间 并且更新left和右边界
			res = append(res, []int{left, right})
			//更新当前的左右边界，并进入下一个循环
			left = intervals[i][0]
			right = intervals[i][1]
		} else {
			// 否则就是重叠的情况了
			// 更新最大右边界，注意这里需要比较最大值才是右边界
			right = max(right, intervals[i][1])
		}
	}
	// 最后一个要手动加入到集合中，不管是冲不重叠都得手动添加
	res = append(res, []int{left, right})
	return res
}
