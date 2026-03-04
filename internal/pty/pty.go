package pty

import (
	"os"
	"os/exec"

	"github.com/creack/pty"
)

func StartShell() (*os.File, error) {
	cmd := exec.Command("bash")
	return pty.Start(cmd)
}
