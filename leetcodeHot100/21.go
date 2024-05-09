package leetcodeHot100

// 合并两个有序的链表，思路其实就是用两个指针分别指向这两个链表，然后进行遍历到新链表，注意边界问题就可以了
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	//新链表的头节点
	var newHead = &ListNode{Val: 0}
	//用来遍历新链表，给新链表赋值
	cur := newHead
	a := list1
	b := list2
	for a != nil && b != nil {
		if a.Val < b.Val {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next
	}
	//出来了说明a或者b有一个为空了
	if a == nil {
		cur.Next = b
	} else {
		cur.Next = a
	}
	//返回新链表，因为新链表的头是虚拟头，就直接返回第二个就行了
	return newHead.Next
}
