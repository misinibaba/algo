package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		start := i + 1
		end := len(nums) - 1

		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if abs(sum - target) < abs(ans - target) {
				ans = sum
			}

			if sum > target {
				end--
			}
			if sum < target {
				start++
			}
			if sum == target {
				fmt.Println(nums[i] ,nums[start] ,nums[end])
				return ans
			}
		}
	}
	return ans
}

func abs (x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func main() {
	res := threeSumClosest([]int{-1,2,1,-4}, 1)
	fmt.Println(res)
}
