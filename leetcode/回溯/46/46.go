package main

import (
	"fmt"
)

var ans [][]int
func permute(nums []int) [][]int {
	ans = [][]int{}
	backTrace(nums, make([]int, 0), make(map[int]bool))
	return ans
}

func backTrace(nums, path []int, used map[int]bool) {
	if len(path) == len(nums) {
		var tem []int
		for _, v := range path {
			tem = append(tem, v)
		}
		ans = append(ans, tem)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[nums[i]] {
			continue
		}
		used[nums[i]] = true
		path = append(path, nums[i])
		backTrace(nums, path, used)
		path = path[:len(path)-1]
		used[nums[i]] = false
	}
}

func main() {
	res := permute([]int{1,2,3})
	fmt.Println(res)
}
