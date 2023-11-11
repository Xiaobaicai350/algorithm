package leetcodeHot100

func searchMatrix(matrix [][]int, target int) bool {
	//进行安全性校验
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	//获取行高和列高
	ls := len(matrix)
	rs := len(matrix[0])
	//初始化遍历二维数组的位置为左下角，这个位置是固定的，因为这个二维数组是有序的，可以进行简化
	l := ls - 1
	r := 0
	for l >= 0 && r < rs {
		if matrix[l][r] > target {
			//如果当前值比目标值大，说明要往上找
			l--
		} else if matrix[l][r] < target {
			//如果当前值比目标值小，说明要往右边找
			r++
		} else { //说明   当前值==目标值 了，这时候就可以返回true了
			return true
		}
	}
	//如果到这里都找不到，说明这个二维数组里面没有target
	return false
}
