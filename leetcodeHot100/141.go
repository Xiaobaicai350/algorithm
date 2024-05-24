package leetcodeHot100

// 本题思想其实就是用一个set，因为set里面存储的是node的地址，
// 如果遍历一遍回来的话，就会找到同样的node，返回true
func hasCycle(head *ListNode) bool {
	set := map[*ListNode]struct{}{}
	for head != nil {
		//这里其实是判断set中是否存在key为head的值，如果存在flag就为true
		_, flag := set[head]
		if flag {
			return true
		}
		set[head] = struct{}{}
		head = head.Next
	}
	return false
}
