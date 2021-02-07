package main

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	maxans := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j] + 1, dp[i])
			}
		}
		maxans = max(dp[i], maxans)
	}
	return maxans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	res := lengthOfLIS([]int{1,3,6,7,9,4,10,5,6})
	fmt.Println(res)
}
