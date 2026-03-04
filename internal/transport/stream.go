package transport

import (
	"io"
)

// Pipe creates a bidirectional stream between two endpoints
func Pipe(a io.ReadWriter, b io.ReadWriter) {
	go io.Copy(a, b)
	io.Copy(b, a)
}
