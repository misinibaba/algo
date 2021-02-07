package main

import "fmt"

func mergeSort(nums []int) []int {
	lens := len(nums)
	if lens < 2 {
		return nums
	}

	mid := lens / 2
	left := nums[:mid]
	right := nums[mid:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}

func main() {
	nums := []int{3,4,2,1,5,5,9,8,7}
	res := mergeSort(nums)
	fmt.Println(res)
}