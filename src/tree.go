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
	recursiveDir(path)

}

func recursiveDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Println("path is not exist")
	}
	files, dirErr := ioutil.ReadDir(dir)
	if dirErr != nil {
		panic(dirErr)
	}
	for _, file := range files {
		if file.IsDir() {
			fmt.Print("d ")
			// recursiveDir(file.Name())
		} else {
			fmt.Print("f ")
		}
		fmt.Println(file.Name())
	}
}
