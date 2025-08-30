package main

import (
	"fmt"
)

func fizzBuzz(n int) []string {
	var answer []string
	if n%3 == 0 && n%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if n%3 == 0 {
		fmt.Println("Fizz")
	} else if n%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println("oops, let's choose another number!")
	}

	return answer
}

func main() {
	fizzBuzz(10)
}
