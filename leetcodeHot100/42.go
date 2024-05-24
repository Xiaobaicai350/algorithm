package leetcodeHot100

/*
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
*/
//这道题其实就是找到左右的最大值，然后依次往中间进行遍历，ans的加是通过两边最大值-遍历节点的值得到的。
func trap(height []int) int {
	l := 0               // l 是左指针，初始指向数组的起始位置
	r := len(height) - 1 // r 是右指针，初始指向数组的最后一个元素
	lMax := height[l]    // lMax 记录从左到当前位置的最高柱子高度
	rMax := height[r]    // rMax 记录从右到当前位置的最高柱子高度
	ans := 0             // rMax 记录从右到当前位置的最高柱子高度
	for l < r {
		if lMax < rMax { //说明该更新lMax了
			ans += lMax - height[l]
			l++
			lMax = max(lMax, height[l])
		} else { //该更新rMax了
			ans += rMax - height[r]
			r--
			rMax = max(rMax, height[r])
		}
	}
	return ans
}
