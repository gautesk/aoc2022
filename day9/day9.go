package main

import (
	"Goworld/lib"
	"fmt"
	"strconv"
	"strings"
)

type Pair struct {
	xPos int
	yPos int
}

func createPair(xPos int, yPos int) Pair {
	pair := Pair{xPos, yPos}
	return pair
}

func updateHead(direction string, currPair Pair) Pair {
	switch direction {
	case "R":
		return createPair(currPair.xPos+1, currPair.yPos)
	case "L":
		return createPair(currPair.xPos-1, currPair.yPos)
	case "U":
		return createPair(currPair.xPos, currPair.yPos+1)
	case "D":
		return createPair(currPair.xPos, currPair.yPos-1)
	}

	fmt.Println("SOMETHIGN WENT VERY WRONG")
	return currPair
}

func updateTail(currH Pair, currT Pair) Pair {
	// Dont move
	diffXPos := currH.xPos - currT.xPos
	diffYPos := currH.yPos - currT.yPos
	if (diffXPos <= 1 && diffXPos >= -1) && (diffYPos <= 1 && diffYPos >= -1) {
		return currT
	}

	// Move up/down
	if diffXPos == 0 {
		if diffYPos > 0 {
			return createPair(currT.xPos, currT.yPos+1)
		} else {
			return createPair(currT.xPos, currT.yPos-1)
		}
	}

	// Move left/right
	if diffYPos == 0 {
		if diffXPos > 0 {
			return createPair(currT.xPos+1, currT.yPos)
		} else {
			return createPair(currT.xPos-1, currT.yPos)
		}
	}

	// Move diagionally
	var newXPos int
	var newYPos int
	if diffYPos > 0 {
		newYPos = currT.yPos + 1
	} else {
		newYPos = currT.yPos - 1
	}

	if diffXPos > 0 {
		newXPos = currT.xPos + 1
	} else {
		newXPos = currT.xPos - 1
	}

	return createPair(newXPos, newYPos)

}

func task9A(input []string) {
	visited := make(map[Pair]bool)
	h := createPair(0, 0)
	t := createPair(0, 0)

	visited[t] = true

	for _, instruction := range input {
		fmt.Println("Current position of h: ", h)

		splitInstructions := strings.Split(instruction, " ")
		direction := splitInstructions[0]
		times, _ := strconv.Atoi(splitInstructions[1])

		for i := 0; i < times; i++ {
			h = updateHead(direction, h)
			t = updateTail(h, t)
			if !visited[t] {
				visited[t] = true
			}
		}
		fmt.Println("Position of head: ", h, " and position of tail ", t)
	}

	fmt.Println("Lengden på besøkte noder er: ", len(visited))
}

func task9B(input []string) {
	visited := make(map[Pair]bool)
	h := createPair(0, 0)

	tailList := [9]Pair{}
	visited[tailList[8]] = true

	for _, instruction := range input {
		splitInstructions := strings.Split(instruction, " ")
		direction := splitInstructions[0]
		times, _ := strconv.Atoi(splitInstructions[1])

		for i := 0; i < times; i++ {
			h = updateHead(direction, h)

			tailList[0] = updateTail(h, tailList[0])

			for i := 1; i < len(tailList); i++ {
				tailList[i] = updateTail(tailList[i-1], tailList[i])
			}

			if !visited[tailList[8]] {
				visited[tailList[8]] = true
			}
		}

	}
	fmt.Println("Lengden på besøkte noder er: ", len(visited))

}

func main() {
	input := lib.ReadAndSplit("day9/day9.txt", "\n")

	task9A(input)
	task9B(input)
}
