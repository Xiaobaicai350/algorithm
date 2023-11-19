package leetcodeHot100

func rotateSlice(nums []int, k int) {
	//这题的技巧性很强
	//翻转前k个到数组尾部
	//1.首先翻转整个数组
	//2.把新数组分成两部分，从第k个元素进行分割
	//3.先翻转前k个
	//4.再翻转后面的

	// 防止数组长度小于k
	k %= len(nums)
	// 翻转整个数组
	reverseSlice(nums, 0, len(nums)-1)
	// 翻转前半部分
	reverseSlice(nums, 0, k-1)
	// 翻转后半部分
	reverseSlice(nums, k, len(nums)-1)
}

// 翻转下标从start到end的数组
func reverseSlice(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
