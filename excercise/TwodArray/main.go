package main

import (
	"fmt"
)

func maximumWealth(accounts [][]int) int {
	var max int = 0
	var tempWealthList []int

	for i := 0; i < len(accounts); i++ {
		var tempWealth int = 0
		for j := 0; j < len(accounts[i]); j++ {
			tempWealth += accounts[i][j]
		}
		tempWealthList = append(tempWealthList, tempWealth)
	}

	for _, wealth := range tempWealthList {
		if wealth > max {
			max = wealth
		}
	}

	fmt.Println(max)
	return max
}

var accounts = [][]int{{1, 5}, {7, 3}, {3, 5}}

func main() {
	maximumWealth(accounts)
}
