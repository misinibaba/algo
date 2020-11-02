package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			ans++
		} else if ans != 0 {
			return ans
		}
	}
	return ans
}

func main() {
	res := lengthOfLastWord("Hello World")
	fmt.Println(res)
}
