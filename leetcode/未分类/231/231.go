package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	return n > 0 && n & (n - 1) == 0
}

func main() {
	res := isPowerOfTwo(13)
	fmt.Println(res)
}
