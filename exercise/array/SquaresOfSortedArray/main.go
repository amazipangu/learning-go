package main

import (
	"fmt"
	"sort"
)

func sortedSquares(nums []int) []int {
	var sortedArray []int

	for _, num := range nums {
		sortedArray = append(sortedArray, num*num)
	}
	sort.Slice(sortedArray, func(i, j int) bool {
		return sortedArray[i] < sortedArray[j]
	})
	return sortedArray
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(nums))
}
