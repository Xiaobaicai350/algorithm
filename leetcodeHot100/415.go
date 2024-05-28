package leetcodeHot100

import "strconv"

/*
输入：num1 = "11", num2 = "123"
输出："134"
*/
func addStrings(num1 string, num2 string) string {
	//保存结果
	ans := ""
	//保存进位信息
	add := 0
	//i1和i2是指向num1和num2末尾的指针
	i1 := len(num1) - 1
	i2 := len(num2) - 1
	for i1 >= 0 || i2 >= 0 || add != 0 {
		x := 0
		y := 0
		if i1 >= 0 {
			x = int(num1[i1] - '0')
		}
		if i2 >= 0 {
			y = int(num2[i2] - '0')
		}
		result := x + y + add
		//把个位上的信息都保存下来，添加到原来字符串的前面
		//strconv.Itoa的作用是把数字转换成字符串
		ans = strconv.Itoa(result%10) + ans
		//保存进位信息
		add = result / 10

		//把i1和i2往前移动
		i1--
		i2--
	}
	return ans
}
