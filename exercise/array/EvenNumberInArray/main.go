package main

import (
	"fmt"
	"strconv"
)

func findNumbers(nums []int) int {
	var numberString string
	var evenNumberDigitsCount int = 0
	var tempArray []int

	for i := 0; i < len(nums); i++ {
		numberString = strconv.Itoa(nums[i])
		if len(numberString)%2 == 0 {
			tempArray = append(tempArray, len(numberString))
		}
		if len(tempArray) > evenNumberDigitsCount {
			evenNumberDigitsCount = len(tempArray)
		}
	}

	return evenNumberDigitsCount
}

func main() {
	nums := []int{12, 345, 2, 6, 7896}
	fmt.Println(findNumbers(nums))
}
