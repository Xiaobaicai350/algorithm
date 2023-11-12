package main

import "fmt"

func main() {
	res := make([]int, 1)
	demo(res)
	fmt.Printf("%d", res)
}
func demo(res []int) {
	res[0] = 1
	res = append(res, 2)
}
