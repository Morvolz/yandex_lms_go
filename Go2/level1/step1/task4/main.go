package main

import (
	"io"
)

func Copy(r io.Reader, w io.Writer, n uint) error {
	buf := make([]byte, n)
	data, err := io.ReadFull(r, buf)

	if err != nil {
		return err
	}

	_, writeErr := w.Write(buf[:data])
	return writeErr
}
