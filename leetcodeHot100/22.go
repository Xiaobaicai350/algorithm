package leetcodeHot100

var result []string

// https://www.bilibili.com/video/BV16K411R7nm/?spm_id_from=333.337.search-card.all.click&vd_source=2259e5459a8cfd21bcf92bc46bf3beda
func generateParenthesis(n int) []string {
	result = make([]string, 0)
	recursive("", 0, 0, n)
	return result
}

// str代表当前递归树的结果path
// l代表当前的左括号的数量
// r代表当前的右括号的数量
// n代表题目给的括号的对数
func recursive(str string, l int, r int, n int) {
	//当左括号数量等于右括号数量的时候，并且都等于题目给的括号的对数的时候，就是要添加啦
	if l == n && r == n {
		result = append(result, str)
	}
	//先让二叉树想左走，就是先加左括号，再加右括号
	if l < n {
		recursive(str+"(", l+1, r, n)
	}
	//并且如果r>l的时候说明不能往下走了，这是条件
	if r < l {
		recursive(str+")", l, r+1, n)
	}
}
