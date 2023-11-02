package leetcodeHot100

// start 代表将要翻转的第一个节点
// end	 代表将要翻转的最后一个节点
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Val: 0,
	}
	dummy.Next = head
	start := dummy
	end := dummy
	for true {
		//先把end往后移动k步，如果不足k步就到达尾部了，就不翻转了
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		//记录startNext和endNext
		startNext := start.Next
		endNext := end.Next
		//将end的next字段制空，方便进行翻转
		end.Next = nil
		start.Next = reverse(start.Next)
		startNext.Next = endNext
		end = startNext
		start = end
	}
	return dummy.Next
}

// 翻转函数，传入head就会将后面的进行翻转
// 返回值是翻转后的头节点（之前的末尾）
func reverse(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
