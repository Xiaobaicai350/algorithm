package leetcodeHot100

func sortList(head *ListNode) *ListNode {
	//特殊情况进行处理
	if head == nil || head.Next == nil {
		return head
	}
	//为了找到中间位置，定义快慢指针
	left := head
	right := head.Next
	for right != nil && right.Next != nil {
		//慢指针走一步
		left = left.Next
		//快指针走两步
		right = right.Next.Next
	}
	//当走到这里的时候，left刚好指向中间位置
	//记录后半部分的指针
	mid := left.Next
	//断开前半部分跟后半部分的连接
	left.Next = nil
	//递归，注意返回的是排序后的头节点
	l := sortList(head)·
	r := sortList(mid)
	res := &ListNode{Val: 0}
	ans := res
	for l != nil && r != nil {
		if l.Val < r.Val {
			res.Next = l
			l = l.Next
		} else {
			res.Next = r
			r = r.Next
		}
		res = res.Next
	}
	if l == nil {
		res.Next = r
	} else {
		res.Next = l
	}
	return ans.Next
}
