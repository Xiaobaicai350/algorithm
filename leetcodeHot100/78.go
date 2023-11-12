package leetcodeHot100

var res [][]int

func subsets(nums []int) [][]int {
	//本题可以抽象出一颗树，向一边走就是不要这个值，向另一边走就是要这个值
	res = [][]int{}
	//从下标为0开始
	recur(0, nums, []int{})
	return res
}

func recur(i int, nums []int, chosen []int) {
	//如果到达树的末尾了，就可以开始返回了
	if i == len(nums) {
		//拷贝临时的结果数组
		tmp := make([]int, len(chosen))
		copy(tmp, chosen)
		res = append(res, tmp)
		return
	}
	//选这个的逻辑，把这个树加入到临时的结果数组中
	chosen = append(chosen, nums[i])
	recur(i+1, nums, chosen)
	//不选这个的逻辑
	chosen = chosen[:len(chosen)-1]
	recur(i+1, nums, chosen)
}
