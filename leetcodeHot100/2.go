package leetcodeHot100

// 不能用转换成int或者long型的数字的方法，因为这种方法会受限于int和long的数据类型的长度，并且只能跑过99%的测试用例。。。。
// 需要下面这种方法

/*
这种方法其实就是模拟了加法进位，然后不用考虑会超过int和long，
因为两个个位数相加不可能大于20，相对于超过int还远呢
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 新建虚拟头结点，用于返回新链表
	var dummyNode = &ListNode{Val: 0}
	//用于遍历新链表
	cur := dummyNode
	//用来保存进位信息
	sum := 0
	for l1 != nil || l2 != nil || sum == 1 { //如果sum==1，但是l1和l2都为空说明进位了
		//至于这里为什么要判断，因为可能是l2为nil或者sum==0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		//新给结果集加节点
		cur.Next = &ListNode{Val: sum % 10}
		//指针后移
		cur = cur.Next
		//查看是否需要进位
		sum /= 10
	}
	return dummyNode.Next
}
