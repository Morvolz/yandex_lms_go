package main

import "io"

type w interface {
	Write(p []byte) (n int, err error)
}

func WriteString(s string, w io.Writer) error {
	_, err := w.Write([]byte(s))
	return err
}

func main() {

}
