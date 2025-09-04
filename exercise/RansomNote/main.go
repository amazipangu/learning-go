package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	magazineRunes := make(map[rune]int)

	for _, char := range magazine {
		magazineRunes[char]++
	}

	for _, c := range ransomNote {
		if magazineRunes[c] == 0 {
			return false
		}
		magazineRunes[c]--
	}
	return true
}

func main() {
	fmt.Println(canConstruct("aaaaabdd", "aaaaabbddd"))
}
