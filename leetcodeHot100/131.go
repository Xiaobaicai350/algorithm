package leetcodeHot100

var (
	pathPartition []string // 放已经回文的子串
	resPartition  [][]string
)

func partition(s string) [][]string {
	pathPartition, resPartition = make([]string, 0), make([][]string, 0)
	dfsPartition(s, 0)
	return resPartition
}

func dfsPartition(s string, start int) {
	if start == len(s) { // 如果起始位置等于s的大小，说明已经找到了一组分割方案了
		tmp := make([]string, len(pathPartition))
		copy(tmp, pathPartition)
		resPartition = append(resPartition, tmp)
		return
	}
	for i := start; i < len(s); i++ {
		str := s[start : i+1]
		if isPalindromePartition(str) { // 是回文子串
			pathPartition = append(pathPartition, str)
			dfsPartition(s, i+1)                                 // 寻找i+1为起始位置的子串
			pathPartition = pathPartition[:len(pathPartition)-1] // 回溯过程，弹出本次已经添加的子串
		}
	}
}

// 验证是否是回文串的方法
func isPalindromePartition(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
