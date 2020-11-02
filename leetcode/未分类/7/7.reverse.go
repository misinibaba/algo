package _

import (
	"fmt"
)

func reverse(x int) int {
	ans := 0
	for x != 0 {
		ans = ans * 10 + x % 10
		if !( -(1<<31) <= ans && ans <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return ans
}

func main() {
	res := reverse(123)
	fmt.Println(res)
}
