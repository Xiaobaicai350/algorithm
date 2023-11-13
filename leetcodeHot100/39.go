package leetcodeHot100

//var (
//	res  [][]int
//	path []int
//)
//
//func combinationSum(candidates []int, target int) [][]int {
//	res, path = make([][]int, 0), make([]int, 0, len(candidates))
//	sort.Ints(candidates) // 排序，为剪枝做准备
//	dfsCombinationSum(candidates, 0, target)
//	return res
//}
//
//func dfsCombinationSum(candidates []int, start int, target int) {
//	if target == 0 { // target 不断减小，如果为0说明达到了目标值
//		tmp := make([]int, len(path))
//		copy(tmp, path)
//		res = append(res, tmp)
//		return
//	}
//	for i := start; i < len(candidates); i++ {
//		if candidates[i] > target { // 剪枝，提前返回
//			break
//		}
//		path = append(path, candidates[i])
//		dfsCombinationSum(candidates, i, target-candidates[i])
//		path = path[:len(path)-1]
//	}
//}
