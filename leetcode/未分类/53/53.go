package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	tmp := nums[0]
	ans := tmp
	for i := 1; i < len(nums); i++ {
		tmp = max(nums[i], tmp + nums[i])
		ans = max(ans, tmp)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	res := maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4})
	fmt.Println(res)
}
