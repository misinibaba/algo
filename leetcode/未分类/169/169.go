package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	major := 0
	count := 0
	for _, v := range nums {
		if count == 0 {
			major = v
		}

		if v == major {
			count++
		} else {
			count--
		}
	}
	return major
}

func main() {
	res := majorityElement([]int{2,2,1,1,1,2,2})
	fmt.Println(res)
}
