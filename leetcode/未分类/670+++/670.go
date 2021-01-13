package main

import (
	"fmt"
)

func maximumSwap(num int) int {

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	res := maximumSwap(2736)
	fmt.Println(res)
}
