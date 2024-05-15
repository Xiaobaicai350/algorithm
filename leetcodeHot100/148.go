package leetcodeHot100

import "sort"

/*
排序以head为头的链表,思路是把链表转换成数组，然后把数组进行排序，然后再转换成新链表
*/
/*
输入：head = [4,2,1,3]
输出：[1,2,3,4]
*/
func sortList(head *ListNode) *ListNode {
	var ints []int
	cur := head
	for cur != nil {
		ints = append(ints, cur.Val)
		cur = cur.Next
	}
	sort.Ints(ints)
	resHead := &ListNode{}
	cur = resHead
	for i := 0; i < len(ints); i++ {
		node := &ListNode{Val: ints[i]}
		cur.Next = node
		cur = cur.Next
	}
	return resHead.Next
}
