package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) (ans [][]int) {
	dfs := func(int, int) {}
	sort.Ints(candidates)
	var path []int
	dfs = func(sum int, index int) {
		if sum == target {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := index; i < len(candidates); i++ {
			if sum + candidates[i] > target {
				return
			}

			path = append(path, candidates[i])
			dfs(candidates[i] + sum, i)
			path = path[:len(path) - 1]
		}
	}
	dfs(0, 0)
	return
}
func main() {
	res := combinationSum([]int{2,3,6,7}, 7)
	fmt.Println(res)
}
