package leetcode

func jump(nums []int) int {
	length := len(nums)
	end := 0
	//用于记录最大可以走过的位置
	maxPosition := 0
	steps := 0
	//注意这里，没有必要访问最后一个元素
	//因为访问最后一个元素之前，我们的边界一定大于等于最后一个位置，否则就无法跳到最后一个位置了
	for i := 0; i < length-1; i++ {
		//其实就是全部遍历，遍历得到最远的位置之后，如果走到了那个位置的话，就记录步数
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}
