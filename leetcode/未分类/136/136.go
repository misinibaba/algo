package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	single := 0
	for _, v := range nums {
		single ^= v
	}
	return single
}

func main() {
	res := singleNumber([]int{2, 2, 3})
	fmt.Println(res)
}
