package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var dfs func(int, int, []int)
	dfs = func(index int, sum int, path []int) {
		if sum == target {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := index; i < len(candidates); i++ {
			if candidates[i] + sum > target {
				break
			}
			path = append(path, candidates[i])
			dfs(i, sum + candidates[i], path)
			path = path[:len(path) - 1]
		}
	}
	dfs(0, 0, make([]int, 0))
	return
}
func main() {
	res := combinationSum([]int{2,3,6,7}, 7)
	fmt.Println(res)
}
