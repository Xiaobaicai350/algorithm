package leetcodeHot100

/*
本题的思想就是把链表中的数据全存放在切片中，
然后用两个指针分别指向数组头和数组尾部，然后一直向中间移动，注意边界问题就好啦
*/
func isPalindrome(head *ListNode) bool {
	list := make([]int, 0)
	//cur用于遍历链表，把链表的元素都添加到数组中，方便从两边进行验证
	cur := head
	for cur != nil {
		list = append(list, cur.Val)
		cur = cur.Next
	}
	//现在list存储着所有链表的数据
	//之后开始用两个指针分别向中间验证
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		if list[i] != list[j] {
			return false
		}
	}
	return true
}
