package main

import "fmt"

var nums []int
func quickSort(left, right int) {
	if left >= right {
		return
	}

	s, e := left, right
	tmp := nums[left]
	for s < e {
		for nums[e] >= tmp && s < e {
			e--
		}
		for nums[s] <= tmp && s < e {
			s++
		}
		nums[s], nums[e] = nums[e], nums[s]
	}
	nums[s], nums[left] = nums[left], nums[s]
	quickSort(left, s - 1)
	quickSort(e + 1, right)
}

func main() {
	nums = []int{3,4,2,1,5,5,9,8,7}
	quickSort(0, len(nums) - 1)
	fmt.Println(nums)
}