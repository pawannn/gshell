package transport

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/pawannn/gshell/internal/pty"
)

func StartListener(port string) error {

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	defer ln.Close()

	fmt.Println("Listening on port", port)

	conn, err := ln.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Println("Client connected from:", conn.RemoteAddr())

	ptmx, err := pty.StartShell()
	if err != nil {
		return err
	}
	defer ptmx.Close()

	// client → shell
	go io.Copy(ptmx, conn)

	// host → shell
	go io.Copy(ptmx, os.Stdin)

	// shell → client + host
	io.Copy(io.MultiWriter(conn, os.Stdout), ptmx)

	return nil
}
