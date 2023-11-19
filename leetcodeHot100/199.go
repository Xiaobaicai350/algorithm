package leetcodeHot100

import "container/list"

// 这个题是102题的变种
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		//如果整棵树都为空，直接返回空数组
		return res
	}
	//创建队列
	queue := list.New()
	//首先先把根节点入队
	queue.PushBack(root)
	for queue.Len() > 0 {
		//保存当前层的长度，然后处理当前层
		//第一次处理的时候因为已经root入队，所以不用特殊处理
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			//和102题不一样的地方在这里
			//只有当在每一层的最后一个元素的时候才加入到返回结果集中，其他的不加
			if i == length-1 {
				res = append(res, node.Val)
			}
		}
	}
	return res
}
