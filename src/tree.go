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
	files, dirErr := ioutil.ReadDir(path)
	if dirErr != nil {
		panic(dirErr)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
