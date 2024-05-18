package leetcodeHot100

func findAnagrams(s string, p string) []int {
	//首先判断s的长度是否大于p的长度，如果len(s)<len(p)，直接返回空数组就可以了,这一步安全性校验是必须的，不加这个过不了
	if len(s) < len(p) {
		return []int{}
	}
	var sCounts [26]int
	var pCounts [26]int
	//初始化两个数组,注意sCounts和pCounts这时赋值长度不一样
	for i := 0; i < len(p)-1; i++ {
		sCounts[s[i]-'a']++
	}
	for i := 0; i < len(p); i++ {
		pCounts[p[i]-'a']++
	}

	var res []int
	//定义滑动窗口
	left := 0
	right := len(p) - 1
	for right < len(s) {
		sCounts[s[right]-'a']++
		//注意这里，go中的数组比较直接就可以用==进行比较，跟Java不一样
		if sCounts == pCounts {
			res = append(res, left)
		}
		sCounts[s[left]-'a']--
		//移动窗口
		left++
		right++
	}
	return res
}
