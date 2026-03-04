package transport

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/pawannn/gshell/internal/auth"
	"github.com/pawannn/gshell/internal/pty"
	"github.com/pawannn/gshell/internal/security"
	"github.com/pawannn/gshell/internal/shutdown"
)

func StartListener(payload security.SessionPayload) error {

	ln, err := net.Listen("tcp", ":"+payload.Port)
	if err != nil {
		return err
	}
	defer ln.Close()

	fmt.Println("Listening on port", payload.Port)

	conn, err := ln.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Println("Client connected from:", conn.RemoteAddr())

	// 🔐 authenticate client
	err = auth.ServerHandshake(conn, payload.Pwd)
	if err != nil {
		fmt.Println("Authentication failed")
		return err
	}

	fmt.Println("Client authenticated")

	ptmx, err := pty.StartShell()
	if err != nil {
		return err
	}
	defer ptmx.Close()

	go shutdown.Wait(func() {
		fmt.Println("\nShutting down session...")
		conn.Close()
		ptmx.Close()
		ln.Close()
	})

	// client → shell
	go io.Copy(ptmx, conn)

	// host → shell
	go io.Copy(ptmx, os.Stdin)

	// shell → client + host
	io.Copy(io.MultiWriter(conn, os.Stdout), ptmx)

	return nil
}
