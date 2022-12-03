package main

import (
	"Goworld/lib"
	"fmt"
)

func task3A(input []string) {
	sum := 0
	for _, x := range input {
		middleIndex := len(x) / 2
		firstCompartment := createMap(x[0:middleIndex])
		for _, v := range []rune(x[middleIndex:]) {
			if firstCompartment[v] {
				sum += runeToAlphabetValue(v)
				break
			}
		}
	}
	fmt.Println(sum)
}

func createMap(content string) map[rune]bool {
	runemap := make(map[rune]bool)
	for _, v := range []rune(content) {
		runemap[v] = true
	}
	return runemap
}

func runeToAlphabetValue(value rune) int {
	valueAsInt := int(value)
	if valueAsInt > 96 {
		return valueAsInt - 96
	}
	return valueAsInt - 38
}

func findGroupBadge(groupList []string) rune {
	firstBackpack := createMap(groupList[0])
	secondBackpack := make(map[rune]bool)
	for _, v := range []rune(groupList[1]) {
		if firstBackpack[v] {
			secondBackpack[v] = true
		}
	}

	for _, v := range []rune(groupList[2]) {
		if secondBackpack[v] {
			return v
		}
	}

	fmt.Println("SOMETHING WENT TERRIBLY WRONG")
	return 0
}

func task3B(input []string) {
	sum := 0
	for i := 0; i < len(input); i += 3 {
		groupList := []string{input[i], input[i+1], input[i+2]}
		groupBadge := findGroupBadge(groupList)
		sum += runeToAlphabetValue(groupBadge)
	}
	fmt.Println(sum)
}

func main() {
	data := lib.ReadAndSplit("day3/day3.txt", "\n")
	task3A(data)
	task3B(data)
}
