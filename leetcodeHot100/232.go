package leetcodeHot100

/*
思路
将一个栈当作输入栈，用于压入push传入的数据；
另一个栈当作输出栈，用于pop和peek操作。
每次 pop 或 peek 时，若输出栈为空则将输入栈的全部数据依次弹出并压入输出栈，这样输出栈从栈顶往栈底的顺序就是队列从队首往队尾的顺序。
*/
// MyQueue 是一个使用两个栈来模拟队列的数据结构。
type MyQueue struct {
	// inStack 用于模拟队列的入队操作。
	inStack []int
	// outStack 用于模拟队列的出队操作。
	outStack []int
}

// Constructor 是 MyQueue 的构造函数，初始化一个空队列。
func Constructor() MyQueue {
	return MyQueue{}
}

// Push 将一个新的元素 x 推入队列的末尾。
func (q *MyQueue) Push(x int) {
	// 将元素 x 推入 inStack。
	q.inStack = append(q.inStack, x)
}

// in2out 是一个辅助函数，用于在需要时将 inStack 的元素移动到 outStack，
// 以便模拟队列的先进先出行为。
func (q *MyQueue) in2out() {
	// 当 inStack 不为空时，循环执行以下操作：
	for len(q.inStack) > 0 {
		// 将 inStack 的栈顶元素（即最后一个元素）推入 outStack。
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		// 移除 inStack 的栈顶元素，即缩小 inStack 的长度。
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}

// Pop 从队列的开头移除并返回元素。这是队列的出队操作。
func (q *MyQueue) Pop() int {
	// 如果 outStack 为空，则需要将 inStack 的元素移动到 outStack。
	if len(q.outStack) == 0 {
		q.in2out()
	}
	// 获取 outStack 的栈顶元素，即队列的队头元素。
	x := q.outStack[len(q.outStack)-1]
	// 移除 outStack 的栈顶元素。
	q.outStack = q.outStack[:len(q.outStack)-1]
	// 返回队列的队头元素。
	return x
}

// Peek 返回队列开头的元素，但不移除它。
func (q *MyQueue) Peek() int {
	// 如果 outStack 为空，则需要将 inStack 的元素移动到 outStack。
	if len(q.outStack) == 0 {
		q.in2out()
	}
	// 返回 outStack 的栈顶元素，即队列的队头元素。
	return q.outStack[len(q.outStack)-1]
}

// Empty 检查队列是否为空。
func (q *MyQueue) Empty() bool {
	// 如果 inStack 和 outStack 都为空，则队列是空的。
	return len(q.inStack) == 0 && len(q.outStack) == 0
}
