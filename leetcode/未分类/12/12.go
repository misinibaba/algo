package main

import "fmt"

func intToRoman(num int) string {
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	intSlice := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	i := 0
	ans := ""
	for i < 13 {
		for num >= intSlice[i] {
			ans += roman[i]
			num -= intSlice[i]
		}
		i++
	}
	return ans
}

func main() {
	res := intToRoman(1994)
	fmt.Println(res)
}
