package leetcodeHot100

func productExceptSelf(nums []int) []int {
	// 计算出数组长度
	length := len(nums)
	// left用来存储除i以外的左边数组之积
	left := make([]int, length)
	// right用来存储除i之外的右边数组之积
	right := make([]int, length)
	// ans用来存储i的左右数组之积
	ans := make([]int, length)
	// 初始化左数组
	left[0] = 1
	// 初始化右数组
	right[length-1] = 1
	// 给左数组赋值
	for i := 1; i < length; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	// 给右数组赋值，右数组是从右往左进行推导的
	for i := length - 2; i > -1; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	// 结果集等于相应的左数组乘相应的右数组
	for i := 0; i < length; i++ {
		ans[i] = left[i] * right[i]
	}
	return ans
}

// [1,2,3,4]
// left ：[1,1,2,6]
// right: [24,12,4,1]
// ans  : [24,12,8,6]
