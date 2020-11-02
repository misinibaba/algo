package main

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				ans[i][j] = 1
			} else {
				ans[i][j] = ans[i - 1][j] + ans[i][j - 1]
			}
		}
	}
	return ans[m - 1][n - 1]
}

func main() {
	res := uniquePaths(7,3)
	fmt.Println(res)
}
