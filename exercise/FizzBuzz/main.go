package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	var answer []string

	for i := 0; i < n; i++ {
		j := i + 1

		if j%3 == 0 && j%5 == 0 {
			answer = append(answer, "FizzBuzz")
		} else if j%3 == 0 {
			answer = append(answer, "Fizz")
		} else if j%5 == 0 {
			answer = append(answer, "Buzz")
		} else {
			answer = append(answer, strconv.Itoa(j))
		}

	}

	fmt.Println(answer) /* [1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz] */
	return answer
}

func main() {
	fizzBuzz(15)
}
