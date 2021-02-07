package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	var dfs func(int, int, []int)
	dfs = func(begin int, sum int, path []int) {
		if sum == target {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := begin; i < len(candidates); i++ {
			if sum + candidates[i] > target {
				break
			}

			if i > begin && candidates[i] == candidates[i - 1] {
				continue
			}

			path = append(path, candidates[i])
			dfs(i + 1, sum + candidates[i], path)
			path = path[:len(path) - 1]
		}
	}
	dfs(0, 0, make([]int, 0))
	return
}

func main() {
	res := combinationSum2([]int{10,1,2,7,6,1,5}, 8)
	fmt.Println(res)
}
