package leetcodeHot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	length := len(lists)
	tail := dummy
	//先把k条链表合成一条链表，之后在进行排序
	for i := 0; i < length; i++ {
		for lists[i] != nil {
			temp := lists[i]
			lists[i] = lists[i].Next
			tail.Next = temp
			tail = tail.Next
		}
	}
	//排序
	list := sortList(dummy.Next)
	return list
}
