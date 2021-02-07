package main

import (
	"fmt"
)

func permuteUnique(nums []int) (ans [][]int)  {
	var dfs func([]int, map[int]bool)
	dfs = func(path []int, used map[int]bool) {
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			if i > 0 && !used[i - 1] && nums[i] == nums[i - 1] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			dfs(path, used)
			path = path[:len(path) - 1]
			used[i] = false
		}
	}
	dfs(make([]int, 0), make(map[int]bool))
	return
}


func main() {
	res := permuteUnique([]int{1,1,2})
	fmt.Println(res)
}
