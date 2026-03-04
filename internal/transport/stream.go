package transport

import (
	"io"
)

// Pipe creates a bidirectional stream between two endpoints
func Pipe(a io.ReadWriter, b io.ReadWriter) {
	go io.Copy(a, b)
	io.Copy(b, a)
}

// FanOut sends data from src to multiple writers
func FanOut(src io.Reader, writers ...io.Writer) {
	buf := make([]byte, 4096)

	for {
		n, err := src.Read(buf)
		if err != nil {
			return
		}

		for _, w := range writers {
			w.Write(buf[:n])
		}
	}
}
