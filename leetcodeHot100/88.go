package leetcodeHot100

func mergeInt(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for p1 != m && p2 != n {
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	if p1 == m {
		sorted = append(sorted, nums2[p2:]...)
	}
	if p2 == n {
		sorted = append(sorted, nums1[p1:]...)
	}
	//合并后数组不应由函数返回，而是存储在数组 nums1 中
	copy(nums1, sorted)
}
