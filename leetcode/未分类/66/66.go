package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
		if carry == 0 {
			break
		}
	}
	if carry  == 1 {
		// digits...代表把digits的元素打散成多个元素插入到第一个slice中
		digits = append([]int{1}, digits...)
	}

	return digits
}

func main() {
	res := plusOne([]int{9})
	fmt.Println(res)
}
