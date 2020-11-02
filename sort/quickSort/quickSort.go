package main

import "fmt"

var nums []int
func quickSort(left, right int) {
	if left > right {
		return
	}

	tmp := 0
	start := left
	end := right

	for start < end {
		tmp = nums[start]
		for nums[start] <= nums[end] && start < end {
			end--
		}
		nums[start] = nums[end]

		for nums[start] < nums[end] && start < end {
			start++
		}
		nums[end] = tmp
	}

	quickSort(left, end - 1)
	quickSort(start + 1, right)
}

func main() {
	nums = []int{1,6,2,1,4,72,5,44,2}
	quickSort(0, len(nums) - 1)
	fmt.Println(nums)
}