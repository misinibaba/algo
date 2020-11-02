package main

import (
	"fmt"
)

func mySqrt(x int) int {
	left := 0
	right := x
	ans := -1
	for left <= right{
		mid := (right + left) / 2
		num := mid * mid

		if num <= x {
			left = mid + 1
			ans = mid
		} else {
			right = mid - 1
		}
	}
	return ans
}

func main() {
	res := mySqrt(1)
	fmt.Println(res)
}
