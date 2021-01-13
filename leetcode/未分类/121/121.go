package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var ans, min int
	min = prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}
		ans = max(prices[i] - min, ans)
	}
	return ans
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func main() {
	res := maxProfit([]int{})
	fmt.Println(res)
}
