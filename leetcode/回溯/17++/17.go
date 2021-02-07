package main

import (
	"fmt"
)

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) (ans []string) {
	var dfs func(int, string)
	dfs = func(index int, path string) {
		if len(path) == len(digits) {
			ans = append(ans, path)
			return
		}

		letter := phoneMap[digits[index:index+1]]
		for _, v := range letter {
			dfs(index + 1, path + string(v))
		}
	}
	dfs(0, "")
	return
}

func main() {
	res := letterCombinations("23")
	fmt.Println(res)
}
