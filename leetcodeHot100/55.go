package leetcodeHot100

func canJump(nums []int) bool {
	//cover表示从当前节点能走到最远的位置
	cover := 0
	for i := 0; i < len(nums); i++ {
		//如果i比cover还大了，说明路径已经断了，走不到终点了
		if i > cover {
			return false
		}
		//更新cover值
		cover = max(cover, i+nums[i])
	}
	return true
}
