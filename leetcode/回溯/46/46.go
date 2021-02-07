package main

import (
	"fmt"
)

func permute(nums []int) (ans [][]int) {
	var dfs func([]int, map[int]bool)
	dfs = func(path []int, used map[int]bool) {
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}

			used[nums[i]] = true
			path = append(path, nums[i])
			dfs(path, used)
			path = path[:len(path) - 1]
			used[nums[i]] = false
		}
	}
	dfs(make([]int, 0), make(map[int]bool))
	return
}


func main() {
	res := permute([]int{1,2,3})
	fmt.Println(res)
}
