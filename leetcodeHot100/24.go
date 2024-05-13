package leetcodeHot100

func swapPairs(head *ListNode) *ListNode {
	//创建虚拟头结点，后面就可以一视同仁了
	dummyNode := &ListNode{
		Val:  0,
		Next: nil,
	}
	dummyNode.Next = head
	//cur用来遍历链表
	cur := dummyNode
	for cur.Next != nil && cur.Next.Next != nil {
		node1 := cur.Next
		node2 := cur.Next.Next
		node3 := cur.Next.Next.Next

		//只需要改变三个指针
		node1.Next = node3
		node2.Next = node1
		cur.Next = node2
		//让cur往后移动两步
		cur = cur.Next.Next
	}
	return dummyNode.Next
}
