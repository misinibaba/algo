package _

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	ans := 0
	copy := x
	for x != 0 {
		ans = ans * 10 + (x % 10)
		x /= 10
	}
	return ans == copy
}

func main() {
	res := isPalindrome(-121)
	fmt.Println(res)
}
