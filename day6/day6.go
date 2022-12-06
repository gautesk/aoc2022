package main

import (
	"Goworld/lib"
	"fmt"
)

func isUnique(section string) bool {
	values := make(map[rune]bool)
	for _, c := range section {
		if values[c] {
			return false
		}
		values[c] = true
	}
	return true
}

func solve(input string, size int) {
	for i := 0; i < len(input); i++ {
		section := input[i : i+size]
		if isUnique(section) {
			fmt.Println(i+size, input[i])
			break
		}
	}
}

func main() {
	input := lib.ReadFile("day6/day6.txt")
	solve(input, 4)
	solve(input, 14)

}
