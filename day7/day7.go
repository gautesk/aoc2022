package main

import (
	"Goworld/lib"
	"fmt"
	"strconv"
	"strings"
)

func buildMap(input []string) map[string]Directory {
	directoryMap := make(map[string]Directory)
	path := "/"
	for _, cmd := range input {
		cmdList := strings.Split(cmd, "\n")
		if len(cmdList) == 1 {
			cdCmd := cmdList[0]
			cdLocation := strings.TrimSpace(strings.Split(cdCmd, " ")[1])

			if cdLocation == ".." {
				if path != "/" {
					pathSplit := strings.Split(path, "/")
					path = strings.Join(pathSplit[0:len(pathSplit)-1], "/")
				}
			} else if cdLocation == "/" {
				path = "/"
			} else {
				path = path + "/" + cdLocation
			}
			path = strings.Replace(path, "//", "/", -1)
		} else {
			var children []string
			fileSizes := 0
			for _, file := range cmdList {
				if file != "" && file != "ls" {
					fileSplit := strings.Split(file, " ")
					if fileSplit[0] == "dir" {
						children = append(children, fileSplit[1])
					} else {
						filesize, _ := strconv.Atoi(fileSplit[0])
						fileSizes += filesize
					}
				}
			}

			updatedDir := UpdateDirectory(fileSizes, children)
			directoryMap[path] = updatedDir

			for _, c := range children {
				if path == "/" {
					directoryMap[path+c] = InitDirectory()
				} else {
					directoryMap[path+"/"+c] = InitDirectory()
				}
			}
		}
	}

	return directoryMap
}

func task6A(input []string) {
	directoryMap := buildMap(input)
	sum := 0
	for key, _ := range directoryMap {
		size := getSize(key, directoryMap)
		if size <= 100000 {
			sum += size
		}
	}
	fmt.Println(sum)
}

func task6B(input []string) {
	directoryMap := buildMap(input)
	TOTAL_SPACE := 70000000
	NEEDED_SPACE := 30000000

	used_space := getSize("/", directoryMap)
	unused_space := TOTAL_SPACE - used_space
	required_freed_up := NEEDED_SPACE - unused_space

	closest_size := 70000000
	for key, _ := range directoryMap {
		size := getSize(key, directoryMap)
		if size > required_freed_up && size < closest_size {
			closest_size = size
		}
	}
	fmt.Println("The closest size is: ", closest_size)
}

func getSize(name string, dirMap map[string]Directory) int {
	activeDir := dirMap[name]
	size := activeDir.filesize
	for _, c := range activeDir.children {
		if name == "/" {
			size += getSize(name+c, dirMap)
		} else {
			size += getSize(name+"/"+c, dirMap)
		}
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
	task6B(trimmedInput)
}
