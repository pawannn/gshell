package pty

import (
	"os"

	"github.com/creack/pty"
)

func Resize(ptmx *os.File, rows int, cols int) error {
	return pty.Setsize(ptmx, &pty.Winsize{
		Rows: uint16(rows),
		Cols: uint16(cols),
	})
}
