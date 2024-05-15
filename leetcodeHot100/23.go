package leetcodeHot100

/*
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
*/

/*
这道题其实就是把链表都合并成一个链表，然后进行排序就可以了
*/
func mergeKLists(lists []*ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := dummyHead
	//合并链表
	for i := 0; i < len(lists); i++ {
		cur.Next = lists[i]
		for cur.Next != nil {
			cur = cur.Next
		}
	}
	//排序链表
	return sortList(dummyHead.Next)
}
