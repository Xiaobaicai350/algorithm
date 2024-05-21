package leetcodeHot100

import "container/list"

// zigzagLevelOrder 实现二叉树的锯齿形层序遍历
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	queue := list.New()
	queue.PushBack(root)
	leftToRight := true // 定义一个标志，表示当前层的遍历方向

	for queue.Len() > 0 {
		//level指的是层级
		levelSize := queue.Len()
		level := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			element := queue.Front()
			queue.Remove(element)
			node := element.Value.(*TreeNode)
			if leftToRight { //正序填充
				level[i] = node.Val
			} else { //倒序填充
				level[levelSize-1-i] = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		result = append(result, level)
		leftToRight = !leftToRight // 反转遍历方向
	}

	return result
}
