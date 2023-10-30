package leetcodeHot100

import "math"

type MinStack struct {
	//存储值的栈
	stack []int
	//辅助栈
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		//定义切片，模拟出栈的效果
		stack: []int{},
		//初始化一个很大很大的值，确保第一个数字一定会比他小
		//这就意味着，minStack里面的值永远比stack多一个
		minStack: []int{math.MaxInt64},
	}
}

// 压栈操作
func (this *MinStack) Push(val int) {
	//把值添加到普通栈中
	this.stack = append(this.stack, val)
	//得到辅助栈栈顶的值作为top
	top := this.minStack[len(this.minStack)-1]
	//比较当前值和辅助栈栈顶的值，得到最小值加入到辅助栈中
	this.minStack = append(this.minStack, min(val, top))
}

// 弹栈操作
func (this *MinStack) Pop() {
	//下面这两行其实是把两个栈都截取最后一个位置的，相当于两个栈都弹栈了
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

// 拿到栈顶的元素值
func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

// 拿到当前栈中的最小值
func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
