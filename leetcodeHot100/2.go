package leetcodeHot100

/*
不能用转换成int或者long型的数字的方法，因为这种方法会受限于int和long的数据类型的长度，并且只能跑过99%的测试用例。。。。
这种方法其实就是模拟了加法进位，然后不用考虑会超过int和long，因为两个个位数相加不可能大于20，相对于超过int还远呢
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//定义虚拟头结点，存储结果
	dummyNode := &ListNode{}
	//用来往后添加结果集节点
	cur := dummyNode
	//保存进位信息
	sum := 0
	//由于sum中保存着进位信息，如果sum为1的时候也可以进行相加
	for l1 != nil || l2 != nil || sum == 1 {
		//注意，这下面不能写成else if的情况，因为这俩都可能不为空
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		node := &ListNode{Val: sum % 10}
		cur.Next = node
		sum /= 10
		cur = cur.Next
	}
	return dummyNode.Next
}
