package leetcodeHot100

// 最终的结果集
var ans [][]int

// 用来存储每一个小组合结果
var path []int

// 标记这个数字是否被使用了
var used []bool

func permute(nums []int) [][]int {
	//注意这三个切片的创建参数
	ans = make([][]int, 0)
	path = make([]int, 0, len(nums))
	used = make([]bool, len(nums))
	dfsPermute(nums, 0)
	return ans
}
func dfsPermute(nums []int, idx int) {
	if idx == len(nums) {
		//这里为什么要新建一个tmp的原因是，如果把path都传入进ans里面，保存的其实是一个地址，所以要传入一个新的对象，防止修改
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
			dfsPermute(nums, idx+1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}

}
