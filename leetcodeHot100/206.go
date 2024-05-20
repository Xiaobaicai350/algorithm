package leetcodeHot100

func reverseList(head *ListNode) *ListNode {
	//fast用于遍历链表
	fast := head
	//slow是fast后面的一个节点
	var slow *ListNode
	for fast != nil {
		next := fast.Next
		//一共需要改变一个指针域就可以了
		fast.Next = slow
		//移动fast和slow
		slow = fast
		fast = next
	}
	//现在fast为nil，slow为新链表的头
	return slow
}
