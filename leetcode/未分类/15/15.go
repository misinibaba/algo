package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums) - 2; i++{
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i - 1] == nums[i] {
			continue
		}

		start := i + 1
		end := len(nums) - 1
		for start < end {
			sum := nums[start] + nums[end] + nums[i]
			if sum > 0 {
				end--
			} else {
				start++
			}

			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[start], nums[end]})
				for nums[start] == nums[start + 1] && start < end {
					start++
				}
				for nums[end] == nums[end - 1] && start < end {
					end--
				}
				start++
				end--
			}
		}
	}
	return ans
}

func main() {
	res := threeSum([]int{-2,0,0,2,2})
	fmt.Println(res)
}
