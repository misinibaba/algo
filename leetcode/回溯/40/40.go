package main

import (
	"fmt"
	"sort"
)

var ans [][]int
func combinationSum2(candidates []int, target int) [][]int {
	ans = [][]int{}
	sort.Ints(candidates)
	backTrace(0, target, 0, make([]int, 0), candidates, make(map[int]bool))
	return ans
}

func backTrace(index, target, sum int, path, candidates []int, used map[int]bool) {
	if sum == target {
		var tem []int
		for _, v := range path {
			tem = append(tem, v)
		}
		ans = append(ans, tem)
		return
	}

	for i := index; i < len(candidates); i++ {
		if sum + candidates[i] > target {
			return
		}
		if i > index && candidates[i] == candidates[i - 1] {
			continue
		}

		path = append(path, candidates[i])
		backTrace(i + 1, target, sum + candidates[i], path, candidates, used)
		path = path[:len(path) - 1]

	}
}

func main() {
	res := combinationSum2([]int{10,1,2,7,6,1,5}, 8)
	fmt.Println(res)
}
