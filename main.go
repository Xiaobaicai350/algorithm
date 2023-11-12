package main

import (
	"algorithm/leetcodeHot100"
	"fmt"
)

func main() {
	ints := []int{1, 2, 3}
	permute := leetcodeHot100.Permute(ints)
	for i, temp := range permute {
		fmt.Println(i, "  ", temp)
	}
}
