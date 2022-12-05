package main

import (
	"Goworld/lib"
	"fmt"
	"strings"
)

func mapXYZ(value string) string {
	if value == "X" {
		return "A"
	} else if value == "Y" {
		return "B"
	}
	return "C"
}

func mapXYZ_pt2(p1 string, p2 string) string {
	allValues := []string{"A", "B", "C"}

	if p2 == "Y" {
		return p1
	}

	shouldWin := p2 == "Z"
	for _, v := range allValues {
		if v != p1 {
			youWinOutcome := youWin(p1, v)
			if youWinOutcome && shouldWin {
				return v
			}

			if !youWinOutcome && !shouldWin {
				return v
			}
		}
	}

	fmt.Println("THIS DID NOT GO WELL")
	return p1
}

func getWeaponScore(value string) int {
	if value == "A" {
		return 1
	} else if value == "B" {
		return 2
	}
	return 3
}

func youWin(p1 string, p2 string) bool {
	if (p1 == "A" && p2 == "B") || (p1 == "B" && p2 == "C") || (p1 == "C" && p2 == "A") {
		return true
	}
	return false
}

func getOutcomeScore(p1 string, p2 string) int {
	if p1 == p2 {
		return 3
	}

	if youWin(p1, p2) {
		return 6
	}
	return 0
}

func task2A(input []string) {
	totalScore := 0

	for _, line := range input {
		lineAsList := strings.Split(line, " ")
		p1 := lineAsList[0]
		p2 := mapXYZ(lineAsList[1])

		totalScore += getWeaponScore(p2)
		totalScore += getOutcomeScore(p1, p2)
	}

	fmt.Println(totalScore)
}

func task2B(input []string) {
	totalScore := 0

	for _, line := range input {
		lineAsList := strings.Split(line, " ")
		p1 := lineAsList[0]
		p2 := mapXYZ_pt2(p1, lineAsList[1])

		totalScore += getWeaponScore(p2)
		totalScore += getOutcomeScore(p1, p2)
	}
	fmt.Println(totalScore)
}

func main() {
	data := lib.ReadAndSplit("day2/day2.txt", "\n")
	task2A(data)
	task2B(data)
}
