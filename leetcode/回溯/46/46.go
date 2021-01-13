package main

import (
	"fmt"
)

func permute(nums []int) (ans [][]int) {
	var dfs func(map[int]bool)
	var path []int
	dfs = func(used map[int]bool) {
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			dfs(used)
			path = path[:len(path) - 1]
			used[i] = false
		}
	}
	dfs(make(map[int]bool))
	return
}


func main() {
	res := permute([]int{1,2,3})
	fmt.Println(res)
}
