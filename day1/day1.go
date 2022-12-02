package main

import (
	"Goworld/lib"
	"fmt"
	"sort"
)

func task1A(input []string) {
	maxWeight := 0
	for _, e := range input {
		sum := 0
		numberList := lib.ParseNumbersAsInt(e)
		for _, x := range numberList {
			sum += x
		}

		if sum > maxWeight {
			maxWeight = sum
		}
	}

	fmt.Println("Task 1A solution is: ", maxWeight)
}

func task1B(input []string) {

	topThreeWeight := []int{0, 0, 0}

	for _, e := range input {
		sum := 0
		numberList := lib.ParseNumbersAsInt(e)
		for _, x := range numberList {
			sum += x
		}

		for index, topValue := range topThreeWeight {
			if sum > topValue {
				topThreeWeight[index] = sum
				break
			}
		}
		sort.Ints(topThreeWeight)
	}

	finalSum := 0
	for _, x := range topThreeWeight {
		finalSum += x
	}

	fmt.Println("Task 1B solution is: ", finalSum)
}

func main() {
	data := lib.ReadAndSplit("day1/day1A.txt", "\n\n")
	task1A(data)
	task1B(data)
}
