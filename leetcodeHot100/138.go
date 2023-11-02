package leetcodeHot100

func copyRandomList(head *Node) *Node {
	//创建map，用来保存原节点和我们复制的节点的对应关系
	hashMap := make(map[*Node]*Node)
	//新建个当前遍历的节点

	cur := head

	for cur != nil {
		//存储原节点和新节点的对应关系
		hashMap[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}
	//返回原链表的头部
	cur = head
	//之后就是给节点的对应关系赋值，比如node中的next和random
	//是基于原链表的关系进行赋值的
	for cur != nil {
		//得到新节点
		newNode := hashMap[cur]
		nextNode := hashMap[cur.Next]
		randomNode := hashMap[cur.Random]
		//开始赋值
		newNode.Next = nextNode
		newNode.Random = randomNode
		//向后遍历原链表
		cur = cur.Next
	}
	//得到新链表的头节点
	return hashMap[head]
}
