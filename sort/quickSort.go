package sort

import "math/rand"

func quickSortMain(arr []int) {
	quickSort(arr)
}

func quickSort(nums []int) {
	// 获取数组的长度。
	n := len(nums)
	// 如果数组长度小于或等于1，无需排序，直接返回。
	if n <= 1 {
		return
	}

	// 随机选择一个索引作为基准值，以优化性能，避免在近乎有序的数组中的性能退化。
	pivotIndex := rand.Intn(n)
	// 将随机选中的基准值交换到数组第一个位置。
	nums[0], nums[pivotIndex] = nums[pivotIndex], nums[0]

	// 选择数组的第一个元素作为基准值。
	pivot := nums[0]
	left, right := -1, n

	for left < right {
		//交换后左右指针需要同时往中间移动一位
		left++
		right--
		for nums[left] < pivot {
			left++
		}
		for nums[right] > pivot {
			right--
		}
		//这时left指向从左到右第一个比基准值大的元素
		//right指向从右到左第一个比基准值小的元素
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	//退出上面的循环后，基准的左边都是比基准小的元素，右边都是比基准大的元素
	//这时left==right
	// 递归地对基准值左边部分的数组进行快速排序。
	quickSort(nums[:right+1])
	// 递归地对基准值右边部分的数组进行快速排序。
	quickSort(nums[right+1:])
}
