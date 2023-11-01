package leetcodeHot100

import "container/list"

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
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
		var tmpArr []int
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmpArr = append(tmpArr, node.Val)
		}
		//将每一层的结果放入结果集
		res = append(res, tmpArr)
		//重新将结果集置为空
		tmpArr = nil
	}
	return res
}
