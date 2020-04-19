package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Tree kek")
	args := os.Args[1:]
	fmt.Println("lentgh", len(args))
	var path string
	if len(args) == 0 {
		cwd, err := os.Getwd()
		if err != nil {
			panic("cwd error ")
		}
		path = cwd
	} else {
		absPath, pathError := filepath.Abs(args[0])
		if pathError != nil {
			log.Fatal(pathError)
		}
		path = absPath
	}
	recursiveDir(path, "", 0)
}

func recursiveDir(dir string, file string, counter int32) {
	intCounter := counter
	fullPath := filepath.Join(dir, file)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		fmt.Println(fullPath, " path is not exist")
		return
	}
	files, dirErr := ioutil.ReadDir(fullPath)
	if dirErr != nil {
		panic(dirErr)
	}
	for _, file := range files {
		if file.IsDir() {
			printRow(intCounter, file.Name())
			recursiveDir(fullPath, file.Name(), intCounter+1)
		} else {
			printRow(intCounter, file.Name())
		}
	}
}

func printRow(counter int32, name string) {
	var i int32
	for i = 0; i < counter; i++ {
		fmt.Print("-")
	}
	fmt.Println(name)
}
