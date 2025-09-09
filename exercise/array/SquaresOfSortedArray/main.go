package main

import (
	"fmt"
)

func sortedSquares(nums []int) []int {
	var sortedArray []int

	for _, num := range nums {
		sortedArray = append(sortedArray, num*num)
	}

	return sortedArray
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}
