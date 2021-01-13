package main

import (
	"fmt"
)

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[0]))
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j > 0 {
				dp[i][j] = dp[i][j - 1] + grid[i][j]
			} else if j == 0 && i > 0 {
				dp[i][j] = dp[i - 1][j] + grid[i][j]
			} else if i == 0 && j == 0 {
				dp[i][j] = grid[0][0]
			} else {
				dp[i][j] = min(dp[i][j - 1], dp[i - 1][j]) + grid[i][j]
			}
		}
	}
	return dp[len(grid) - 1][len(grid[0]) - 1]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	res := minPathSum([][]int{{1,3,1}, {1,5,1}, {4,2,1}})
	fmt.Println(res)
}
