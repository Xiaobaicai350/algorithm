package leetcodeHot100

import "sort"

func groupAnagrams(strs []string) [][]string {
	//map中	key为排序后的字符串，value为一个string类型的list
	mp := map[string][]string{}
	for _, str := range strs {
		//得到每个字符串
		s := []byte(str)
		//对字符串进行排序
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		//再转换成字符串
		sortedStr := string(s)
		//添加到key为排序后的字符串的value的list中
		mp[sortedStr] = append(mp[sortedStr], str)
	}
	//创建结果集，长度为mp的大小
	ans := make([][]string, 0, len(mp))
	//遍历map，看能组合成几个list
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
