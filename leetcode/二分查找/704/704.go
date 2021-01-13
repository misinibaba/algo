package main

import "fmt"

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := l + (r - l) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{-1,0,3,5,9,12}
	target := 9
	res := search(nums, target)
	fmt.Println(res)
}
