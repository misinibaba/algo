package main

import (
	"fmt"
)

func moveZeroes(nums []int)  {
	fast := 0
	slow := 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		fast++
	}
	fmt.Println(nums)
}

func main() {
	nums := []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
}
