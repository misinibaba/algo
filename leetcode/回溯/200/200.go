package main

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				dfs(i, j, grid)
				count++
			}
		}
	}
	return count
}

func dfs(x, y int, grid [][]byte) {
	if x >= len(grid) || y >= len(grid[0]) || x < 0 || y < 0 || grid[x][y] == 0 {
		return
	}

	grid[x][y] = 0
	dfs(x + 1, y, grid)
	dfs(x - 1, y, grid)
	dfs(x, y + 1, grid)
	dfs(x, y - 1, grid)

}

func main() {
	input := [][]byte{{1,1,1,1,0}, {1,1,0,1,0}, {1,1,0,0,0}, {0,0,0,0,0}}
	res := numIslands(input)
	fmt.Println(res)
}
