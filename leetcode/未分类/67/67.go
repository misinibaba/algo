package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	ans := ""
	var carry uint8
	for i >= 0 || j >= 0{
		var x uint8
		if i >= 0 {
			x = a[i] - '0'
		}

		var y uint8
		if j >= 0 {
			y = b[j] - '0'
		}

		sum := x + y + carry
		carry = sum / 2
		ans = strconv.Itoa(int(sum % 2)) + ans

		i--
		j--

	}

	if carry > 0 {
		ans = "1" + ans
	}

	return ans
}

func main() {
	res := addBinary("11", "1")
	fmt.Println(res)
}
