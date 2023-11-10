package leetcodeHot100

// go的list模拟栈https://blog.csdn.net/SaberJYang/article/details/125219350
func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	stack := []int{}
	for i := 0; i < length; i++ {
		temperature := temperatures[i]
		//如果栈里有元素的话 并且 当前的元素大于栈顶的元素
		//就需要出栈
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			//得到栈顶元素的值
			prevIndex := stack[len(stack)-1]

			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return ans
}
