package leetcodeHot100

func detectCycle(head *ListNode) *ListNode {
	//其实跟上道题一样，就是一个是返回布尔值一个是返回节点
	set := map[*ListNode]struct{}{}
	for head != nil {
		_, flag := set[head]
		if flag {
			return head
		} else {
			set[head] = struct{}{}
		}
		head = head.Next
	}
	return nil
}
