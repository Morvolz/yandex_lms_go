package main

import (
	"fmt"
	"os"
)

func ReadContent(filename string) string {
	data, _ := os.ReadFile(filename)
	return string(data)
}

func main() {
	fmt.Printf(ReadContent("literature.txt"), "\n")
}
