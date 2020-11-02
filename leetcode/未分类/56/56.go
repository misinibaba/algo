package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i,j int) bool{
		return intervals[i][0] < intervals[j][0]
	})

	if len(intervals) == 0 {
		return make([][]int, 0)
	}

	ans := make([][]int, 0)
	ans = append(ans, intervals[0])
	j := 0
	for i := 1; i < len(intervals); i++ {
		if ans[j][1] >= intervals[i][0] {
			ans[j][1] = max(intervals[i][1], ans[j][1])
		} else {
			j++
			ans = append(ans, intervals[i])
		}
	}
	return ans
}

func max (x,y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	res := merge([][]int{{1,4}, {2,3}})
	fmt.Println(res)
}
