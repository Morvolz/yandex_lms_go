package main

import (
	"bufio"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	f, _ := os.Open(inputFilename)
	fileScanner := bufio.NewScanner(f)
	var i int
	for fileScanner.Scan() {
		if i == lineNum {
			return fileScanner.Text()
		}
		i++
	}
	return ""
}

// func main() {
// 	fmt.Print(LineByNum("task2/literature.txt", 5), "\n")
// }
