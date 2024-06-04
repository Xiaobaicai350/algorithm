package sort

func heapSortMain(arr []int) {
	heapSort(arr)
}

func heapSort(nums []int) []int {
	end := len(nums) - 1 // 获取数组的最后一个元素的索引
	// 构建最大堆
	for i := end / 2; i >= 0; i-- {
		sink(nums, i, end) // 调整元素i，确保堆的性质
	}
	// 逐步将最大的元素放到数组的末尾，并进行堆调整
	for i := end; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0] // 交换堆顶元素和当前末尾元素
		end--                               // 缩小堆的大小
		sink(nums, 0, end)                  // 对新的堆顶元素进行下沉操作，维护堆的性质
	}
	return nums
}

// sink 是一个辅助函数，用于在堆中下沉一个元素，确保子堆满足最大堆的性质。
func sink(heap []int, root, end int) {
	for {
		child := root*2 + 1 // 计算根节点的左子节点
		if child > end {    // 如果子节点索引超出堆的范围，则结束
			return
		}
		// 如果存在右子节点，并且右子节点的值大于左子节点的值，则选择右子节点
		if child < end && heap[child] <= heap[child+1] {
			child++
		}
		// 如果根节点的值已经大于子节点的值，则不需要继续下沉
		if heap[root] > heap[child] {
			return
		}
		// 交换根节点和较大的子节点的值
		heap[root], heap[child] = heap[child], heap[root]
		root = child // 更新根节点为子节点的索引，继续下一轮比较
	}
}
