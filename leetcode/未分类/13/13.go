package main

import "fmt"

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	pre := 0
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := m[s[i]]
		if cur >= pre {
			ans += cur
		} else {
			ans -= cur
		}
		pre = cur
	}
	return ans
}

func main() {
	res := romanToInt("III")
	fmt.Println(res)
}
