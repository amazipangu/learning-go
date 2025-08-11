package main

import (
	"fmt"
)

// Declare an int variable "i" with the value 20
// Assign "i" to a floating-point value named "f"
// Print out "i" and "f"
func main() {
	var i int = 20
	var f float64 = float64(i)

	fmt.Println(i, f)
}
