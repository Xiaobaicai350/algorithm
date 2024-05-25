package leetcodeHot100

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []rune
	for _, ch := range s {
		if pair, ok := pairs[ch]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != pair {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}
