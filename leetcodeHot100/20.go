package leetcodeHot100

/*
输入：s = "()[]{}"
输出：true
*/
func isValid(s string) bool {
	n := len(s)
	//这段代码是提高效率的，没有这段代码也可以
	if n%2 == 1 {
		return false
	}
	//保存括号的对应关系
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	//模拟栈的实现
	var stack []byte
	for i := 0; i < len(s); i++ {
		char := s[i]
		if pair, ok := pairs[char]; ok {
			//如果现在栈里为空，并且当前为)]}，肯定是不合理的情况
			//或者是现在栈顶元素不等于当前的，也是不合理的情况
			if len(stack) == 0 || stack[len(stack)-1] != pair {
				return false
			}
			//弹栈
			stack = stack[:len(stack)-1]
		} else {
			//压栈
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
