package main

import (
	"fmt"
)

func findNumbers(nums []int) int {
	var evenNumberDigitsCount int = 0

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		var digitsCount int = 0

		for n > 0 {
			n /= 10
			digitsCount++
		}

		if digitsCount%2 == 0 {
			evenNumberDigitsCount++
		}

	}

	return evenNumberDigitsCount
}

func main() {
	nums := []int{12, 345, 2, 6, 7896}
	fmt.Println(findNumbers(nums))
}
