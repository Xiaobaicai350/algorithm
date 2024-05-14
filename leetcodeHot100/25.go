package leetcodeHot100

/*
输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyNode := &ListNode{}
	dummyNode.Next = head
	// start 代表将要翻转的第一个节点的原始前一个节点
	// end	 代表将要翻转的最后一个节点
	start := dummyNode
	end := dummyNode
	for true {
		//先把end往后移动k位，如果不足k位就不需要移动了，说明要退出循环了
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		//从上面循环跳出来需要先判断end是否为nil，如果end为nil，说明没必要翻转了
		if end == nil {
			break
		}
		//记录下原始头节点
		startNext := start.Next
		//记录下原始尾节点的后一个节点
		endNext := end.Next
		end.Next = nil
		start.Next = reverse(startNext)
		startNext.Next = endNext
		//更新一下start和end，方便下次翻转
		end = startNext
		start = startNext
	}
	return dummyNode.Next
}

// 翻转以node为头的链表
// 返回值是翻转后的链表头结点
func reverse(node *ListNode) *ListNode {
	cur := node
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
