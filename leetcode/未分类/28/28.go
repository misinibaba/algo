package main

import "fmt"

func strStr(haystack string, needle string) int {
	l1 := len(haystack)
	l2 := len(needle)

	if l1 < l2 {
		return -1
	}
	if l1 == 0 {
		return 0
	}

	for i := 0; i < l1 - l2; i++ {
		if haystack[i:i+l2] == needle {
			return i
		}
	}
	return -1
}

func main() {
	res := strStr("a", "a")
	fmt.Println(res)
}
