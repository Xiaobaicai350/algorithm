package leetcodeHot100

// 最终的结果集
var ans [][]int

// 用来存储每一个小组合结果
var path []int

// 标记这个数字是否被使用了
var used []bool

func permute(nums []int) [][]int {
	ans = make([][]int, 0)
	path = make([]int, 0, len(nums))
	used = make([]bool, len(nums))
	dfs_premute(nums, 0)
	return ans
}
func dfs_premute(nums []int, idx int) {
	if idx == len(nums) {
		//这里不知道为什么要新建一个来存放path，和其他语言不太一样
		tmp := make([]int, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {
		//如果这个数字没被用过，就使用一下
		if !used[i] {
			path = append(path, nums[i])
			used[i] = true
			dfs_premute(nums, idx+1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

}