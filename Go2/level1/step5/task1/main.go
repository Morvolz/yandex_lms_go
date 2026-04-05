package main

import (
	"bytes"
	"context"
	"io"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	data := []byte{}
	chunk := make([]byte, 2048)

	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()

		default:
		}
		n, err := r.Read(chunk)

		if err != nil {
			if err == io.EOF {
				return false, nil
			} else {
				return false, err
			}
		}

		data = append(data, chunk[:n]...)
		if bytes.Index(data, seq) >= 0 {
			return true, nil
		}

		if len(data) >= len(seq) {
			data = data[len(data)-len(seq):]
		}
	}
}
