package main

import "fmt"

func isValid(s string) bool {
	if len(s) % 2 == 1 {
		return false
	}

	arr := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if arr[s[i]] > 0 {
			if len(stack) == 0 || arr[s[i]] != stack[len(stack) - 1] {
				return false
			}
			stack = stack[:len(stack) - 1]
		} else {
			stack = append(stack, s[i])
		}
	}

	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {
	res := isValid("){")
	fmt.Println(res)
}
