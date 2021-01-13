package main

import (
	"fmt"
)

func subsets(nums []int) (ans [][]int) {
	var dfs func(int)
	var path []int
	dfs = func(begin int) {
		ans = append(ans, append([]int{}, path...))
		if len(path) == len(nums) {
			return
		}

		for i := begin; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path) - 1]
		}
	}
	dfs(0)
	return
}


func main() {
	res := subsets([]int{1,2,3})
	fmt.Println(res)
}
