package main

import "fmt"

func removeElement(nums []int, val int) int {
	first, end := 0, 0
	for end < len(nums) {
		if nums[end] == val {
			end++
		} else {
			nums[first] = nums[end]
			first++
			end++
		}
	}
	nums = nums[:first]
	return len(nums)
}

func main() {
	nums := []int{0,1,2,2,3,0,4,2}
	res := removeElement(nums, 2)
	fmt.Println(res)
}
