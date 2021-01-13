package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	var newS string
	for i := 0; i < len(s); i++ {
		if check(s[i]) {
			newS += string(s[i])
		}
	}

	newS = strings.ToLower(newS)
	l := 0
	r := len(newS) - 1
	for l < r {
		lChar := newS[l]
		rChar := newS[r]

		if lChar == rChar {
			l++
			r--
			continue
		} else {
			return false
		}
	}
	return true
}

func check(s byte) bool {
	return s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z' || s >= '0' && s <= '9'
}

func main() {
	res := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(res)
}
