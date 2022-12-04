package main

import (
	"Goworld/lib"
	"fmt"
	"strconv"
	"strings"
)

func task4A(input []string) {
	count := 0
	for _, line := range input {
		splitLine := strings.Split(line, ",")
		e1 := strings.Split(splitLine[0], "-")
		e2 := strings.Split(splitLine[1], "-")

		e1Min, _ := strconv.Atoi(e1[0])
		e1Max, _ := strconv.Atoi(e1[1])
		e2Min, _ := strconv.Atoi(e2[0])
		e2Max, _ := strconv.Atoi(e2[1])

		if (e1Min >= e2Min && e1Max <= e2Max) || (e2Min >= e1Min && e2Max <= e1Max) {
			count += 1
		}
	}
	fmt.Println(count)

}

func hasOverlap(e1 []string, e2 []string) bool {
	e1Min, _ := strconv.Atoi(e1[0])
	e1Max, _ := strconv.Atoi(e1[1])
	e2Min, _ := strconv.Atoi(e2[0])
	e2Max, _ := strconv.Atoi(e2[1])

	if e1Max >= e2Min && e1Min < e2Min { // Overlap from below
		return true
	} else if e1Min <= e2Max && e1Max > e2Max { // Overlap from above
		return true
	} else if e1Min >= e2Min && e1Max <= e2Max { // Contained overlap
		return true
	}
	return false
}

func task4B(input []string) {
	count := 0
	for _, line := range input {
		splitLine := strings.Split(line, ",")
		e1 := strings.Split(splitLine[0], "-")
		e2 := strings.Split(splitLine[1], "-")
		if hasOverlap(e1, e2) || hasOverlap(e2, e1) {
			count += 1
		}
	}
	fmt.Println(count)
}

func main() {
	data := lib.ReadAndSplit("day4/day4.txt", "\n")
	task4A(data)
	task4B(data)
}
