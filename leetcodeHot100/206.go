package leetcodeHot100

func reverseList(head *ListNode) *ListNode {
	fast := head
	var slow *ListNode
	for fast != nil {
		//先保存fast后面那个节点
		fastNext := fast.Next
		fast.Next = slow
		slow = fast
		fast = fastNext
	}
	//slow不为空，现在fast已经为空了
	return slow
}
