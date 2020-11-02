package main

import (
	"fmt"
)

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}

			if i == 0  {
				grid[i][j] += grid[i][j - 1]
			} else if j == 0 {
				grid[i][j] += grid[i - 1][j]
			} else {
				grid[i][j] += min(grid[i - 1][j], grid[i][j - 1])
			}
		}
	}
	return grid[m - 1][n - 1]
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
