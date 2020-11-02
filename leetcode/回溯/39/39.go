package main

import (
	"fmt"
	"sort"
)

var ans [][]int
func combinationSum(candidates []int, target int) [][]int {
	ans = [][]int{}
	sort.Ints(candidates)
	backTrace(0, target, 0, make([]int, 0), candidates)
	return ans
}

func backTrace(index, target, sum int, path, candidates []int) {
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

		path = append(path, candidates[i])
		backTrace(i, target, sum + candidates[i], path, candidates)
		path = path[:len(path) - 1]
	}
}

func main() {
	res := combinationSum([]int{3,4,7,8}, 11)
	fmt.Println(res)
}
