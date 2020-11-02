package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums) - 2;i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		start := i + 1
		end := len(nums) - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if sum < 0 {
				start++
			}
			if sum > 0 {
				end--
			}

			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[start], nums[end]})
				// 跳过相同的
				for nums[start] == nums[start + 1] && start + 1 < end {
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
	res := threeSum([]int{0,0,0})
	fmt.Println(res)
}
