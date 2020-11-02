package main

import "fmt"
var ans []string
func generateParenthesis(n int) []string {
	ans = []string{}
	backTrace(n, n, "")
	return ans
}

func backTrace(left, right int, path string) {
	if left == 0 && right == 0 {
		ans = append(ans, path)
		return
	}

	if left > right {
		return
	}

	if left > 0 {
		backTrace(left - 1, right, path + "(")
	}

	if right > 0 {
		backTrace(left, right - 1, path + ")")
	}
}



func main() {
	res := generateParenthesis(1)
	fmt.Println(res)
}
