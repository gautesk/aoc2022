package lib

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadFile(filename string) string {
	content, _ := os.ReadFile(filename)
	return string(content)
}

func ParseNumbers(input string) []string {
	return regexp.MustCompile(`[-+]?\d*\.?\d*`).FindAllString(input, -1)
}

func ParseNumbersAsInt(input string) []int {
	numbersAsString := ParseNumbers(input)

	var numbers []int
	for _, e := range numbersAsString {
		value, _ := strconv.Atoi(e)
		numbers = append(numbers, value)
	}

	return numbers
}

func NumbersFromFile(filename string) []string {
	filecontent := ReadFile(filename)
	return ParseNumbers(filecontent)
}

func NumbersAsIntFromFile(filename string) []int {
	filecontent := ReadFile(filename)
	return ParseNumbersAsInt(filecontent)
}

func ReadAndSplit(filename string, separator string) []string {
	filecontent := ReadFile(filename)
	return strings.Split(filecontent, separator)
}

func ReadDataStrings(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
