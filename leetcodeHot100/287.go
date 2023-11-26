package leetcodeHot100

// 跟142题很像，因为模拟构造了一个环形链表
func findDuplicate(nums []int) int {
	//思路：
	//		数组中的每个元素num指向下一个元素nums[num]
	//		如果有重复数字，肯定存在一个环形链表
	//		环入口节点的前一个节点即为重复的元素
	fast := nums[0]
	slow := nums[0]
	//因为元素都是【0-n】的，所以直接以数组下标构建就可以了
	fast = nums[nums[fast]]
	slow = nums[slow]
	for fast != slow {
		fast = nums[nums[fast]]
		slow = nums[slow]
	}
	fast = nums[0]
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	return slow
}
