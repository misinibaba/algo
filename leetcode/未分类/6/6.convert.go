package _

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	n := 2 * numRows - 2
	for i, char := range s {
		x := i % n
		rows[min(x, n - x)] += string(char)
	}
	return strings.Join(rows, "")
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	s := "LEETCODEISHIRING"
	numRows := 3
	res := convert(s, numRows)
	fmt.Println(res)
}
