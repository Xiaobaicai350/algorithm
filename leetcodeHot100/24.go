package leetcodeHot100

func swapPairs(head *ListNode) *ListNode {
	//创建虚拟头节点
	dummyNode := &ListNode{Val: 0}
	dummyNode.Next = head

	//cur用来遍历链表
	cur := dummyNode
	for cur.Next != nil && cur.Next.Next != nil {
		temp1 := cur.Next
		temp2 := cur.Next.Next
		temp3 := cur.Next.Next.Next

		//这里不能用temp来表示cur.Next因为cur.Next已经被修改了，而temp还是原来的地方
		//这里可能有点难理解，用笔画个图就好了
		cur.Next = temp2
		cur.Next.Next = temp1
		cur.Next.Next.Next = temp3
		//之后再向后移动两步就可以了
		cur = cur.Next.Next
	}
	return dummyNode.Next
}
