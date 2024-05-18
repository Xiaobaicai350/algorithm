package leetcodeHot100

func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]struct{})
	left := 0
	right := 0
	res := 0
	//注意这里的退出条件是right超过界限,因为right是右指针，只有他可能超出界限，如果他超出界限之后才能找到最大值
	for left < len(s) && right < len(s) {
		if _, ok := set[s[right]]; !ok { //如果set中不存在s[right]，说明可以让窗口扩大
			set[s[right]] = struct{}{}
			right++
		} else { //如果存在，就缩小窗口
			delete(set, s[left])
			left++
		}
		res = max(res, right-left)
	}
	return res
}
