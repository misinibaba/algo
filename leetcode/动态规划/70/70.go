package main

import (
	"fmt"
)

func climbStairs(n int) int {
	ans := []int{1, 1}
	for i := 2; i <= n; i++ {
		ans = append(ans, ans[i - 1] + ans[i - 2])
	}
	return ans[n]
}

func main() {
	res := climbStairs(3)
	fmt.Println(res)
}
