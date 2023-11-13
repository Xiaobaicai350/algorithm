package leetcodeHot100

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			//进行遍历，得到直到得到最终的结果，如果dfs有一个返回的是true，说明这个数组里面就有word
			dfs := dfsExist(board, word, i, j, 0)
			if dfs {
				//只要有一个dfs返回是true，就返回就好了
				return true
			}
		}
	}
	//如果遍历完了还没有,说明这个数组里面没有word
	return false
}

func dfsExist(board [][]byte, word string, i int, j int, k int) bool {
	//首先进行边界判断,判断i，j是否合法，访问的这个格子是否在board范围内
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 {
		//如果有一个不满足说明就不在格子里面，直接返回false，说明这个格子不可以作为结果。
		return false
	}
	if board[i][j] != word[k] {
		//如果访问的这个格子的值不等于我们要访问的这个字母，也返回false
		return false
	}
	//如果走到这里了，说明这个是我们要找的第k位字母
	//但是我们要判断一下这个字母是不是我们最终要找的字母的最后一位
	if k == len(word)-1 {
		return true
	}
	//为了防止他被多次访问，我们给他个特殊值，不过不用担心之后数据被破坏，如果后面步骤找不到合适的，会回滚
	board[i][j] = '$'
	//将我们要找的字母位数加一
	k++
	//		上									左									下									右
	res := dfsExist(board, word, i, j-1, k) || dfsExist(board, word, i-1, j, k) || dfsExist(board, word, i, j+1, k) || dfsExist(board, word, i+1, j, k)
	//回滚
	board[i][j] = word[k-1]
	//如果res为true，说明可以找到
	return res
}
