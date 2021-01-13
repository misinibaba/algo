package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = 1 // 左边第一个乘1，由于最左边一个不乘自己，所以不用管
	// ans存入nums每个数字的左边累乘
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i - 1] * nums[i - 1]
	}

	R := 1 // 右边第一个也是乘1
	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] = ans[i] * R
		R *= nums[i] // 这一步用于保存Nums的值，用于下一次往左的累乘
	}
	return ans
}

func main() {
	res := productExceptSelf([]int{1,2,3,4})
	fmt.Println(res)
}
