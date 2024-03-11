package leetcodeHot100

func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		//存储有向图
		edges = make([][]int, numCourses)
		//存储每个节点的入度数量
		indeg  = make([]int, numCourses)
		result []int
	)

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		indeg[info[0]]++
	}
	//创建队列
	q := []int{}
	for i := 0; i < numCourses; i++ {
		//第一次先把为0的入队
		if indeg[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		result = append(result, u)
		for _, v := range edges[u] {
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(result) == numCourses
}
