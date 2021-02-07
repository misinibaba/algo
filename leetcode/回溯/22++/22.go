package main

import "fmt"

func generateParenthesis(n int) (ans []string) {
	var dfs func(string, int, int)
	dfs = func(path string, l int, r int) {
		if len(path) == n * 2{
			ans = append(ans, path)
			return
		}

		if l < n {
			dfs(path + "(", l + 1, r)
		}

		if r < l {
			dfs(path + ")", l, r + 1)
		}
	}
	dfs("", 0, 0)
	return
}



func main() {
	res := generateParenthesis(3)
	fmt.Println(res)
}
