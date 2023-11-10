package leetcodeHot100

func hasCycle(head *ListNode) bool {
	//本题思想其实就是用一个set，因为set里面存储的是node的地址，如果遍历一遍回来的话，就会找到同样的node，返回true
	//这里需要注意的是go语言里面的set的具体表示，还是得学习呀。。。

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
