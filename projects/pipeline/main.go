package main

import (
	"fmt"
)

// реализовать removeDuplicates(in, out chan string)
func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)

	go removeDuplicates(inputStream, outputStream)

	go func() {
		inputStream <- "a"
		inputStream <- "a"
		inputStream <- "b"
		inputStream <- "b"
		inputStream <- "c"
		close(inputStream)
	}()

	for x := range outputStream {
		fmt.Print(x)
	}
	fmt.Print("\n")
}

func removeDuplicates(inputStream chan string, outputStream chan string) {
	var prevStr string
	var curStr string

	for value := range inputStream {
		curStr = value
		if curStr != prevStr {
			outputStream <- curStr
			prevStr = curStr
		}
	}
	close(outputStream)
}
