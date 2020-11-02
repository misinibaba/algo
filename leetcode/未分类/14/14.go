package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]);i++ {
		for j := 1; j < len(strs) ; j++ {
			if i == len(strs[j]) || strs[0][i] != strs[j][i]  {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main() {
	res := longestCommonPrefix([]string{"aa", "a"})
	fmt.Println(res)
}
