package main

import "fmt"

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		}

		// 如果没翻转过
		if nums[left] <= nums[mid] {
			// 如果target在left->mid之间，是则根据二分定位，否则定位到mid->end之间去
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 翻转过，则查看是否在mid->end之间，是的话定位到mid->end，否则定位到left->mid之间
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}
	return -1
}

func main() {
	res := search([]int{5,1,3}, 3)
	fmt.Println(res)
}
