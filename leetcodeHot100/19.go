package leetcodeHot100

/*
首先先定义虚拟头结点，然后接上原来的头结点
之后再定义快慢指针，让fast指针先走n+1步，这样的话fast为null的时候slow刚好为删除节点前面的那个位置
最后再让fast和slow指针同时移动，当fast==nil的时候，slow刚好指向要删除的那个节点，然后进行删除就可以了
*/
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
