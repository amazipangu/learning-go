package main

func numberOfSteps(num int) int {
	var steps int = 0

	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		steps++
	}
	return steps
}

func main() {
	numberOfSteps(1)
}
