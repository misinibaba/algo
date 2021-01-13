package main

import (
	"fmt"
)

func canWinNim(n int) bool {
	return n % 4 == 0
}

func main() {
	res := canWinNim(4)
	fmt.Println(res)
}
