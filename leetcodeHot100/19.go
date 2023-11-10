package leetcodeHot100

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//首先定义虚拟头节点，这样的话就不用担心删除的是头节点了，原因是因为n是从倒着数的，无法判断啥时候是头节点
	dummyHead := &ListNode{Val: 0}
	//给虚拟头节点接上
	dummyHead.Next = head
	//定义快慢指针
	fast := dummyHead
	slow := dummyHead
	//先让fast指针走n+1步，这样的话fast为null的时候slow刚好为删除节点前面的那个位置
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	//然后把fast走到nil，此时slow刚好为删除节点前面的那个位置
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	//进行删除操作
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
