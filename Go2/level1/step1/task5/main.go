package main

import (
	"bytes"
	"io"
)

func Contains(r io.Reader, seq []byte) (bool, error) {
	buffer := make([]byte, 4096)

	_, err := r.Read(buffer)

	if err != nil {
		return false, err
	}

	idx := bytes.Index(buffer, seq)

	if idx == -1 {
		return false, nil
	} else {
		return true, nil
	}
}
