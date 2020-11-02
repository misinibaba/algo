package main

import (
	"fmt"
	"sort"
)

var ans [][]int
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	ans = [][]int{}
	backTrace(nums, make([]int, 0), make(map[int]bool, 0))
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
		if used[i] {
			continue
		}
		// 未使用过used{i - 1}说明是循环过来的，并且跟i一样，说明已经使用过了，就跳过
		if i > 0 && nums[i] == nums[i - 1] && !used[i - 1] {

			continue
		}
		used[i] = true
		path = append(path, nums[i])
		backTrace(nums, path, used)
		path = path[:len(path)-1]
		used[i] = false
	}
}

func main() {
	res := permuteUnique([]int{3,3,0,3})
	fmt.Println(res)
}
