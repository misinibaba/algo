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

}

func main() {
	res := letterCombinations("23")
	fmt.Println(res)
}
