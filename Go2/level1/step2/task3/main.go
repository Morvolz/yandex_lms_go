package main

import (
	"os"
)

func CopyFilePart(inputFilename, outFileName string, startpos int) error {
	inp, err := os.Open(inputFilename)
	inp.Seek(int64(startpos), 0)
	buffer := make([]byte, 8096)
	n, err := inp.Read(buffer)

	out, err := os.Create(outFileName)
	defer func() {
		err = out.Close()
	}()

	_, err = out.Write(buffer[:n])

	return err
}

// func main() {
// 	fmt.Print(CopyFilePart("task3/literature.txt", "task3/literature+n.txt", 10), "\n")
// }
