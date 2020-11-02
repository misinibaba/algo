package main

import "fmt"

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	strLen := len(s)
	max := 0
	start, end := 0, 0
	for i := 0; i < strLen; i++ {
		l1, r1 := getMaxLong(s, i, i)
		l2, r2 := getMaxLong(s, i, i + 1)

		if r2 - l2 > r1 - l1 {
			if r2 - l2 > max {
				start = l2
				end = r2
				max = r2 - l2
			}
		} else {
			if r1 - l1 > max {
				start = l1
				end = r1
				max = r1 - l1
			}
		}

	}
	return s[start:end+1]
}

func getMaxLong(str string, left, right int) (int, int) {
	for left >= 0 && right < len(str) && str[left] == str[right] {
		left--
		right++
	}
	return left + 1, right - 1
}

func main() {
	s := "bb"
	res := longestPalindrome(s)
	fmt.Println(res)
}
