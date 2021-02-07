package main

import (
	"fmt"
)

func movingCount(m int, n int, k int) int {
	var dfs func(int, int) (count int)
	visited := make(map[string]bool)

	dfs = func(x, y int)  (count int) {
		if x > n - 1 || y > m - 1 || getSum(x,y) > k {
			return 0
		}

		if _, exist := visited[fmt.Sprintf("%d,%d", x, y)]; exist {
			return 0
		}

		visited[fmt.Sprintf("%d,%d", x, y)] = true
		count++
		count += dfs(x + 1, y)
		count += dfs(x, y + 1)

		return count
	}
	return dfs(0, 0)
}

func getSum(x, y int) int {
	var sum int
	for x > 0 {
		sum += x % 10
		x /= 10
	}

	for y > 0 {
		sum += y % 10
		y /= 10
	}
	return sum
}

func main() {
	res := movingCount(3, 1, 0)
	fmt.Println(res)
}
