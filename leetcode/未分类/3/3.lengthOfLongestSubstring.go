package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	subStr := ""
	max := 0
	strLen := len(s)
	pos := 0
	for i := 0; i < strLen; i++ {
		pos = strings.IndexAny(subStr, s[i:i+1])
		if pos != -1 {
			subStr = subStr[pos+1:] + s[i:i+1]
		} else {
			subStr += s[i:i+1]
		}

		if len(subStr) > max {
			max = len(subStr)
		}
	}
	return max
}

func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}
