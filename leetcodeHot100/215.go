package leetcodeHot100

import "container/heap"

// 这道题就是模拟实现最小堆，利用最小堆的特性来做题
// 最小堆的堆顶存的是存放元素的最小值
// 我们把n-k个元素都弹出去，数组里只存储k个元素，那么堆顶的就是我们要的那个值
type MinHeap []int

// 返回堆里有多少个元素
func (h MinHeap) Len() int { return len(h) }

// 返回二者中比较小的值
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

// 交换两者的下标
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	//弹出数组尾部的元素
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)
	//先把元素全填进去
	for _, num := range nums {
		heap.Push(h, num)
	}
	//再只保留k个
	for i := 0; i < len(nums)-k; i++ {
		heap.Pop(h)
	}
	return heap.Pop(h).(int)
}
