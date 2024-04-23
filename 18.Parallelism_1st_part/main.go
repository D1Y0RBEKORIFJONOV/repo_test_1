package main

import (
	"fmt"
	"os"
)

func ReadFileGO(filename string, data chan string) {
	str, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	data <- string(str)
}

func main() {
	DATA_INSIDE_FILE := make(chan string)
	go ReadFileGO("file1.txt", DATA_INSIDE_FILE)
	go ReadFileGO("file2.txt", DATA_INSIDE_FILE)
	go ReadFileGO("file3.txt", DATA_INSIDE_FILE)
	go ReadFileGO("file4.txt", DATA_INSIDE_FILE)
	
	for i := 0; i < 4; i++ {
		fmt.Println(<-DATA_INSIDE_FILE)
	}
}
