package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	var ans int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i - 1] {
			ans += prices[i] - prices[i - 1]
		}
	}
	return ans
}
func main() {
	res := maxProfit([]int{7,6,4,3,1})
	fmt.Println(res)
}
