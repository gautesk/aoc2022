package main

import (
	"Goworld/lib"
	"fmt"
	"strconv"
	"strings"
)

func task6A(input []string) {
	directoryMap := make(map[string]Directory)
	marker := "/"

	directoryMap[marker] = InitDirectory("")

	for _, cmd := range input {
		cmdList := strings.Split(cmd, "\n")
		if len(cmdList) == 1 {
			cdCmd := cmdList[0]
			cdLocation := strings.Split(cdCmd, " ")[1]
			if cdLocation == ".." {
				marker = directoryMap[marker].parent
			} else {
				marker = cdLocation
			}
		} else {
			var children []string
			var filesizes int
			for _, file := range cmdList {
				if file != "" && file != "ls" {
					fileSplit := strings.Split(file, " ")
					if fileSplit[0] == "dir" {
						children = append(children, fileSplit[1])
					} else {
						filesize, _ := strconv.Atoi(fileSplit[0])
						filesizes += filesize
					}
				}
			}

			if val, ok := directoryMap[marker]; ok {
				updatedDir := UpdateDirectory(filesizes, val.parent, children)
				directoryMap[marker] = updatedDir
			} else {
				fmt.Println("WE GOT A BIG PROBLEM")
			}

			for _, c := range children {
				if val, ok := directoryMap[c]; ok {
					updatedDir := UpdateDirectory(val.filesize, marker, val.children)
					directoryMap[c] = updatedDir
				} else {
					newDir := InitDirectory(marker)
					directoryMap[c] = newDir
				}
			}
		}
	}

	var answerDirs []SimpleDirectory
	sum := 0
	for key, dir := range directoryMap {
		size := getSize(dir, directoryMap)
		if size < 100000 {
			answerDirs = append(answerDirs, CreateSimpleDirectory(key, size))
			sum += size
		}
	}

	fmt.Println(sum)
}

func getSize(dir Directory, dirMap map[string]Directory) int {
	size := dir.filesize
	for _, c := range dir.children {
		size += getSize(dirMap[c], dirMap)
	}
	return size
}

func main() {
	input := lib.ReadAndSplit("day7/day7.txt", "$")

	var trimmedInput []string
	for _, line := range input {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			trimmedInput = append(trimmedInput, trimmedLine)
		}
	}

	task6A(trimmedInput)
}
