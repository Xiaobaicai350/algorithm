package leetcodeHot100

func partitionLabels(s string) (partition []int) {
	//初始化数组用来存储每一个字符最后出现的位置
	lastPos := [26]int{}
	for i, c := range s {
		lastPos[c-'a'] = i
	}
	start, end := 0, 0
	for i, c := range s {
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}
		//如果到了end节点，说明已经形成一个区间了，可以增加成结果集了
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return
}
