package main

import (
	"Goworld/lib"
	"fmt"
)

func solveTask(numbers []int) {
	for _, number := range numbers {
		for _, number2 := range numbers {
			if number+number2 == 2020 {
				fmt.Println(number, number2, number+number2, number*number2)
				return
			}
		}
	}
}

func main() {

	dog := lib.NumbersAsIntFromFile("template/input_task1.txt")

	solveTask(dog)

	//fileAsString := readFile("template/input_task1.txt")
	//parseNumbers(fileAsString)

	//allStrings := readDataStrings("template/input_small.txt")
	//fmt.Println(allStrings)

	//solveTask(allNumbers)
}
