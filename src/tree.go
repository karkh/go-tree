package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Printf("hello world!")
	args := os.Args[1:]
	fmt.Println(args)
	files, err := ioutil.ReadDir("C:\\")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
