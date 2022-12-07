package main

type Directory struct {
	filesize int
	parent   string
	children []string
}

func UpdateDirectory(filesize int, parent string, children []string) Directory {
	directory := Directory{filesize, parent, children}
	return directory
}

func InitDirectory(parent string) Directory {
	directory := Directory{0, parent, []string{}}
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
