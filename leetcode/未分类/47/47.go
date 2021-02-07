package main

import (
	"fmt"
)

func maxValue(grid [][]int) int {
	for i := 1; i < len(grid); i++ {
		grid[i][0] += grid[i - 1][0]
	}

	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] += grid[0][j - 1]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] += max(grid[i - 1][j], grid[i][j - 1])
		}
	}

	return grid[len(grid) - 1][len(grid[0]) - 1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	res := maxValue([][]int{{1,3,1}, {1,5,1}, {4,2,1}})
	fmt.Println(res)
}
