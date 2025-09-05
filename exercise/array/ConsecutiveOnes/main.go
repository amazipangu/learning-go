/*
Input: nums = [1,1,0,1,1,1]
Output: 3
Explanation:
- The first two digits or the last three digits are consecutive 1s.
- The maximum number of consecutive 1s is 3.
*/

package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	var result int = 0
	var currentCount int = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			currentCount++
		} else {
			currentCount = 0
		}
		if currentCount > result {
			result = currentCount
		}
	}
	return result
}

func main() {
	var nums []int = []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(nums))
}
