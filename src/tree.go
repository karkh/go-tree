package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Tree kek")
	args := os.Args[1:]
	path, pathError := filepath.Abs(args[0])
	if pathError != nil {
		panic(pathError)
	}
	recursiveDir(path, "", 0)

}

func recursiveDir(dir string, path string, counter int32) {
	intCounter := counter
	fullPath := filepath.Join(dir, path)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		fmt.Println("path is not exist")
	}
	files, dirErr := ioutil.ReadDir(fullPath)
	if dirErr != nil {
		panic(dirErr)
	}
	for _, file := range files {
		if file.IsDir() {
			recursiveDir(fullPath, file.Name(), intCounter+1)
		}
		var i int32
		for i = 0; i < intCounter; i++ {
			fmt.Print("-")
		}
		fmt.Println(file.Name())
	}
}
