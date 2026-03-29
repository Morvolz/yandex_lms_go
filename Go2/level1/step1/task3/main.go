package main

import "strings"

type UpperWriter struct {
	UpperString string
}

func (uw *UpperWriter) Write(p []byte) (n int, err error) {
	uw.UpperString = strings.ToUpper(string(p))
	return len(p), nil
}

func main() {

}
