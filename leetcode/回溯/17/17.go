package main

import (
	"fmt"
)
var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
var combinations []string
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack (digits string, index int, path string) {
	if len(path) == len(digits) {
		combinations = append(combinations, path)
	} else {
		letters := phoneMap[digits[index:index+1]]
		for _, v := range letters {
			backtrack(digits, index + 1, path + string(v))
		}
	}
}

func main() {
	res := letterCombinations("23")
	fmt.Println(res)
}
