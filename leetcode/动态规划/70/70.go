package main

import (
	"fmt"
)

func climbStairs(n int) int {
	dp := []int{1, 2}
	for i := 2; i < n; i++ {
		dp = append(dp, dp[i - 1] + dp[i - 2])
	}
	return dp[n - 1]
}

func main() {
	res := climbStairs(3)
	fmt.Println(res)
}
