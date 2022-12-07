package main

type Directory struct {
	filesize int
	children []string
}

func UpdateDirectory(filesize int, children []string) Directory {
	directory := Directory{filesize, children}
	return directory
}

func InitDirectory() Directory {
	directory := Directory{0, []string{}}
	return directory
}
