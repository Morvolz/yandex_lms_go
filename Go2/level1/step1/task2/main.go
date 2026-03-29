package main

import (
	"io"
)

// type r interface {
// 	Read(p []byte) (n int, err error)
// }

func ReadString(r io.Reader) (string, error) {
	s, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(s), err
}
func main() {

}
