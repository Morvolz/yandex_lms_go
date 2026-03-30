package main

import (
	"bufio"
	"fmt"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	f, _ := os.Open(inputFilename)
	fileScanner := bufio.NewScanner(f)
	var i int
	for fileScanner.Scan() {
		i++
		// fmt.Print(fileScanner.Text(), "\n")
		if i == lineNum {
			return fileScanner.Text()
		}
	}
	return ""
}
func main() {
	fmt.Print(LineByNum("task2/literature.txt", 5), "\n")
}
