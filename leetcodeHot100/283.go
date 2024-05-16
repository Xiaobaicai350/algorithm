package leetcodeHot100

/*
原理其实就是定义两个双指针，并且这道题需要保证非零元素的相对顺序，所以双指针指向的是一前一后
*/
func moveZeroes(nums []int) {
	left := 0
	//注意这个循环条件首先得判断长度，再去找是否等于0，防止数组越界,因为可能出现没0的情况
	for left < len(nums) && nums[left] != 0 {
		left++
	}
	//跳出上个循环的 这里left指向的是第一个0
	right := left + 1
	//right从left的下一个开始，也就是从第一个0的后面一个位置开始，主要是为了换位置
	//之后就是right找到非0元素跟left进行交换，如果是0的话就直接跳过，直到right到了数组末尾
	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
