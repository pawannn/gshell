package transport

import (
	"fmt"
	"io"
	"net"
	"os"

	"github.com/pawannn/gshell/internal/auth"
	"github.com/pawannn/gshell/internal/security"
	"github.com/pawannn/gshell/internal/shutdown"
	"golang.org/x/term"
)

func Connect(payload security.SessionPayload) error {

	address := payload.IP + ":" + payload.Port
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	err = auth.ClientHandshake(conn, payload.Pwd)
	if err != nil {
		return err
	}

	fmt.Println("Connected to", payload.Session)

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	go shutdown.Wait(func() {
		fmt.Println("\nClosing connection...")
		conn.Close()
		term.Restore(int(os.Stdin.Fd()), oldState)
	})

	stdio := struct {
		io.Reader
		io.Writer
	}{
		os.Stdin,
		os.Stdout,
	}

	Pipe(conn, stdio)

	return nil
}
