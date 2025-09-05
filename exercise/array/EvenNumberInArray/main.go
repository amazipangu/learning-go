package main

import (
	"fmt"
)

func findNumbers(nums []int) int {
	var evenCount int = 0

	for _, num := range nums {
		var digitsCount int = 0

		for num > 0 {
			num /= 10
			digitsCount++
		}
		if digitsCount%2 == 0 {
			evenCount++
		}
	}

	return evenCount
}

func main() {
	nums := []int{12, 345, 2, 6, 7896}
	fmt.Println(findNumbers(nums))
}
