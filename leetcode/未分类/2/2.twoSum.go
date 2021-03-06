package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for k, v := range nums {
		complement := target - v
		if val, ok := numMap[complement]; ok {
			return []int{val, k}
		}
		numMap[v] = k
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}
