package main

import (
	"Goworld/lib"
	"fmt"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func createContainers(input string) []Stack {
	lines := strings.Split(input, "\n")
	lastLine := lines[len(lines)-1]
	containerLines := lines[:len(lines)-1]
	height := len(containerLines)

	containerNumbers := strings.Split(lastLine, " ")
	width, _ := strconv.Atoi(containerNumbers[len(containerNumbers)-1])

	containerMap := make([]Stack, width)
	for i := 0; i > width; i++ {
		var stack Stack
		containerMap[i] = (stack)
	}

	for i := height - 1; i >= 0; i-- {
		replacedString := strings.Replace(containerLines[i], "    ", "[1]", -1)
		replacer := strings.NewReplacer("[", "", "]", "", " ", "")
		containerList := replacer.Replace(replacedString)
		fmt.Println(containerList)
		for j, c := range containerList {
			if c != 49 {
				containerMap[j].Push(strconv.QuoteRune(c))
			}
		}
	}

	return containerMap
}

func task5A(containers []Stack, operations []string) {
	for _, op := range operations {
		opList := strings.Split(op, " ")
		toMove, _ := strconv.Atoi(opList[1])
		from, _ := strconv.Atoi(opList[3])
		to, _ := strconv.Atoi(opList[5])

		for i := 0; i < toMove; i++ {
			movingC, _ := containers[from-1].Pop()
			containers[to-1].Push(movingC)
		}
	}

	answer := ""
	for _, x := range containers {
		popped, _ := x.Pop()
		answer += popped
	}

	fmt.Println(answer)

}

func task5B(containers []Stack, operations []string) {
	for _, op := range operations {
		opList := strings.Split(op, " ")
		toMove, _ := strconv.Atoi(opList[1])
		from, _ := strconv.Atoi(opList[3])
		to, _ := strconv.Atoi(opList[5])

		var tempStack Stack
		for i := 0; i < toMove; i++ {
			movingC, _ := containers[from-1].Pop()
			tempStack.Push(movingC)
		}

		for !tempStack.IsEmpty() {
			tempC, _ := tempStack.Pop()
			containers[to-1].Push(tempC)
		}
	}

	answer := ""
	for _, x := range containers {
		popped, _ := x.Pop()
		answer += popped
	}
	fmt.Println(answer)
}

func main() {
	input := lib.ReadAndSplit("day5/day5.txt", "\n\n")
	containers := createContainers(input[0])
	task5A(containers, strings.Split(input[1], "\n"))
	task5B(containers, strings.Split(input[1], "\n"))

}
