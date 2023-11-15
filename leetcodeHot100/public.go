package leetcodeHot100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ListNode struct {
	Val  int
	Next *ListNode
}
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func BubbleSort(arr []int, start int, end int) []int {
	for i := start; i < end-1; i++ {
		for j := start; j < end-i-1+start; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
