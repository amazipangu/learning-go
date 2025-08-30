package main

import (
	"fmt"
)

func sum(num1 int, num2 int) int {
	var result int = 0
	if -100 <= num1 && num2 <= 100 {
		result = num1 + num2
	} else {
		fmt.Println("Error")
	}
	return result
}
