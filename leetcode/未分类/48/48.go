package main

import (
	"fmt"
)
var matrix [][]int
func rotate(matrix [][]int)  {
	tmp := 0
	pos1 := 0
	pos2 := len(matrix) - 1
	for pos1 < pos2 {
		add := 0
		for add < pos2 - pos1 {
			tmp = matrix[pos1][pos1 + add]
			matrix[pos1][pos1 + add] = matrix[pos2 - add][pos1]
			matrix[pos2 - add][pos1] = matrix[pos2][pos2 - add]
			matrix[pos2][pos2 - add] = matrix[pos1 + add][pos2]
			matrix[pos1 + add][pos2] = tmp
			add++
		}
		pos1++
		pos2--
	}
}

func main() {
	matrix = [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
