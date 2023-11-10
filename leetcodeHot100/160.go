package leetcodeHot100

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//本题的基本思想是：链表1的长度+链表2的长度=链表2的长度+链表1的长度
	//这样的话在走互相的路的话，肯定会碰面的

	//分别给他俩分配一个指针用来走
	a := headA
	b := headB
	//当a和b相遇的时候就会退出循环啦
	for !(a == b) {
		//说明a自己的那一趟走完了，该走b的了
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		//说明b自己的那一趟走完了，该走a的了
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	//这里返回a和b都可以啦，因为他们俩这时候就是相等的
	return a
}
