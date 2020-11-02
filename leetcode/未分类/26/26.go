package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	first, end := 0, 0
	for end < len(nums) {
		if nums[first] != nums[end] {
			first++
			nums[first] = nums[end]
			end++
		} else {
			end++
		}
	}
	nums = nums[:first+1]
	return len(nums)
}

func main() {
	res := removeDuplicates([]int{})
	fmt.Println(res)
}
