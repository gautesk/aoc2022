package main

import (
	"Goworld/lib"
	"fmt"
	"strconv"
	"strings"
)

const MATRIX_SIZE = 99

func buildMatrix(input []string) [MATRIX_SIZE][MATRIX_SIZE]int {
	matrix := [MATRIX_SIZE][MATRIX_SIZE]int{}
	for i, line := range input {
		lineSplit := strings.Split(line, "")
		for j, v := range lineSplit {
			intVal, _ := strconv.Atoi(v)
			matrix[i][j] = intVal
		}
	}
	return matrix
}

func calculateSight(posY int, posX int, matrix [MATRIX_SIZE][MATRIX_SIZE]int) int {
	selectedTree := matrix[posY][posX]
	score := 4

	// Upwards
	for i := posY - 1; i >= 0; i-- {
		if matrix[i][posX] >= selectedTree {
			score = score - 1
			break
		}
	}

	// Downwards
	for i := posY + 1; i < MATRIX_SIZE; i++ {
		if matrix[i][posX] >= selectedTree {
			score = score - 1
			break
		}
	}

	// Left
	for i := posX - 1; i >= 0; i-- {
		if matrix[posY][i] >= selectedTree {
			score = score - 1
			break
		}
	}
	// Right
	for i := posX + 1; i < MATRIX_SIZE; i++ {
		if matrix[posY][i] >= selectedTree {
			score = score - 1
			break
		}
	}

	if score > 0 {
		return 1
	} else {
		return 0
	}
}

func calculateViewLength(posY int, posX int, matrix [MATRIX_SIZE][MATRIX_SIZE]int) int {
	selectedTree := matrix[posY][posX]
	view := [4]int{1, 1, 1, 1}

	// Upwards
	for i := posY - 1; i > 0; i-- {
		if matrix[i][posX] >= selectedTree {
			break
		} else {
			view[0] += 1
		}
	}

	// Downwards
	for i := posY + 1; i < MATRIX_SIZE-1; i++ {
		if matrix[i][posX] >= selectedTree {
			break
		} else {
			view[1] += 1
		}
	}

	// Left
	for i := posX - 1; i > 0; i-- {
		if matrix[posY][i] >= selectedTree {
			break
		} else {
			view[2] += 1
		}
	}
	// Right
	for i := posX + 1; i < MATRIX_SIZE-1; i++ {
		if matrix[posY][i] >= selectedTree {
			break
		} else {
			view[3] += 1
		}
	}

	return view[0] * view[1] * view[2] * view[3]
}

func calculateInnerVisible(matrix [MATRIX_SIZE][MATRIX_SIZE]int) int {
	sights := 0
	for i := 1; i < MATRIX_SIZE-1; i++ {
		for j := 1; j < MATRIX_SIZE-1; j++ {
			sights += calculateSight(i, j, matrix)
		}
	}
	return sights
}

func getBestView(matrix [MATRIX_SIZE][MATRIX_SIZE]int) int {
	bestScore := 0
	for i := 1; i < MATRIX_SIZE-1; i++ {
		for j := 1; j < MATRIX_SIZE-1; j++ {
			score := calculateViewLength(i, j, matrix)
			if score > bestScore {
				bestScore = score
			}
		}
	}
	return bestScore
}

func task8B(input []string) {
	width := len(input[0])
	height := len(input)
	fmt.Println(width, height)

	matrix := buildMatrix(input)
	bestView := getBestView(matrix)
	fmt.Println("The one with the best view has value: ", bestView)
}

func task8A(input []string) {
	width := len(input[0])
	height := len(input)
	fmt.Println(width, height)

	outerVisible := (MATRIX_SIZE * 4) - 4

	matrix := buildMatrix(input)

	innerVisible := calculateInnerVisible(matrix)
	fmt.Println(outerVisible+innerVisible, outerVisible, innerVisible)

}

func main() {

	input := lib.ReadAndSplit("day8/day8.txt", "\n")

	task8B(input)
}
