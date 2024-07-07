package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	ShowHidden bool
	HiddenTree bool
)

func main() {
	folderFlag := flag.Bool("f", false, "Only show folders")
	showHiddenFlag := flag.Bool("a", false, "Show hidden folders")

	flag.Parse()
	ShowHidden = *showHiddenFlag

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if *folderFlag {
		treeOnlyFolder(cwd, 0)
	} else {
		tree(cwd, 0, true)
	}
}

func tree(path string, depth int, showFiles bool) {
	listDir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(listDir); i++ {
		if listDir[i].IsDir() {
			hidden := _printFolder(listDir[i], depth)
			tree(path+"/"+listDir[i].Name(), depth+1, hidden)
		} else if showFiles {
			_printFile(listDir[i], depth)
		}
	}
}

func treeOnlyFolder(path string, depth int) {
	listDir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(listDir); i++ {
		if !listDir[i].IsDir() {
			continue
		} else {
			_printFolder(listDir[i], depth)
			treeOnlyFolder(path+"/"+listDir[i].Name(), depth+1)
		}
	}
}

func _printFile(file os.DirEntry, depth int) {
	var fileF string
	if !ShowHidden {
		if file.Name()[0] == '.' {
			fileF = "break"
		}
	}
	if fileF != "break" {
		for i := 0; i < depth; i++ {
			fileF += "  "
		}
		fileF += "┕ 📰 " + file.Name()

		fmt.Println(fileF)
	}
}

func _printFolder(folder os.DirEntry, depth int) bool {
	var folderF string
	if !ShowHidden {
		if folder.Name()[0] == '.' {
			folderF = "break"
			return false
		}
	}
	if folderF != "break" {
		for i := 0; i < depth; i++ {
			folderF += "  "
		}
		folderF += "┕ 📁 " + folder.Name()

		fmt.Println(folderF)
		return true
	}
	return false
}
