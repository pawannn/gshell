package transport

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/pawannn/gshell/internal/pty"
	"github.com/pawannn/gshell/internal/security"
)

func StartListener(payload security.SessionPayload) error {

	ln, err := net.Listen("tcp", ":"+payload.Port)
	if err != nil {
		return err
	}
	defer ln.Close()

	fmt.Println("Listening on port", payload.Port)

	expiryDuration := time.Until(time.Unix(payload.Expiry, 0))

	if expiryDuration <= 0 {
		return fmt.Errorf("session already expired")
	}

	connChan := make(chan net.Conn)
	errChan := make(chan error)

	go func() {
		conn, err := ln.Accept()
		if err != nil {
			errChan <- err
			return
		}
		connChan <- conn
	}()

	select {

	case conn := <-connChan:
		defer conn.Close()
		fmt.Println("Client connected from:", conn.RemoteAddr())

		ptmx, err := pty.StartShell()
		if err != nil {
			return err
		}
		defer ptmx.Close()
		go io.Copy(ptmx, conn)
		go io.Copy(ptmx, os.Stdin)
		io.Copy(io.MultiWriter(conn, os.Stdout), ptmx)

	case err := <-errChan:
		return err

	case <-time.After(expiryDuration):
		fmt.Println("Session expired. No clients joined.")
		return nil
	}

	return nil
}
