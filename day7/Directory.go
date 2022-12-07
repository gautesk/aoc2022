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

type SimpleDirectory struct {
	name      string
	totalSize int
}

func CreateSimpleDirectory(name string, totalSize int) SimpleDirectory {
	dir := SimpleDirectory{name, totalSize}
	return dir
}
