package main

func runningSum(nums []int) []int {
	var num int = 0
	for i, v := range nums {
		num += v
		nums[i] = num
	}
	return nums
}
