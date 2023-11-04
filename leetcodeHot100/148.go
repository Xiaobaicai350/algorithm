package leetcodeHot100

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Val: 0, Next: head}
	n := 0
	for node := head; node != nil; node = node.Next {
		n++
	}

	for i := 0; i < n-1; i++ {
		// 使用两个指针prev和curr来遍历链表并比较相邻节点的值
		prev := dummy
		curr := dummy.Next
		for j := 0; j < n-i-1; j++ {
			if curr.Val > curr.Next.Val {
				// 交换相邻节点的值
				temp := curr.Next
				curr.Next = curr.Next.Next
				temp.Next = curr
				prev.Next = temp
			}
			prev = curr
			curr = curr.Next
		}
	}

	return dummy.Next
}
