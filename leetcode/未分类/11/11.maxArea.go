package main

import "fmt"

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max := 0
	for left != right {
		area := min(height[left], height[right]) * (right - left)
		if max < area {
			max = area
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return max
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	res := maxArea([]int{1,8,6,2,5,4,8,3,7})
	fmt.Println(res)
}
