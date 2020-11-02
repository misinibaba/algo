package _

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	ans := 0
	abs := 1
	for i, v := range str {
		if v >= '0' && v <= '9' {
			ans = ans * 10 + int(v - '0')
			if ans > math.MaxInt32 {
				if abs == -1 {
					return math.MinInt32
				}
				return math.MaxInt32
			}
		} else if v == '-' && i == 0 {
			abs = -1
		}  else if v == '+' && i == 0 {
			abs = 1
		} else {
			break
		}
	}
	return ans * abs
}

func main() {
	res := myAtoi("-2147483647")
	fmt.Println(res)
}
